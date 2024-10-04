package publish

import (
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/toolsascode/helm-s3-publisher/internal/git"
	"github.com/toolsascode/helm-s3-publisher/internal/helm"
	"github.com/toolsascode/helm-s3-publisher/internal/helpers"
	"github.com/toolsascode/helm-s3-publisher/internal/plugins"
)

var reportPublish = []Report{}

/*
	helm-s3-publisher REPO [CHART PATHS] [flags]
*/
// Run
func (c *Commands) Run() {

	if err := git.CheckIntall(); err != nil {
		log.Fatalln("git.CheckIntall()", err)
	}

	if err := helm.CheckIntall(); err != nil {
		log.Fatalln("helm.CheckIntall()", err)
	}
	if err := plugins.S3CheckIntall(); err != nil {
		log.Fatalln("plugins.S3CheckIntall()", err)
	}

	var (
		listPath = viper.GetStringSlice("chart.paths")
	)

	log.Infoln("Starting chart list processing...")
	c.chartList(listPath)
	c.GenerateReport(reportPublish)

}

func (c *Commands) chartList(paths []string) {

	var list []string

	if !viper.GetBool("git.lsTree") {
		log.Infoln("Git LS Tree is not enabled!!!")
		list = paths
	} else {
		list = c.gitLsTree(paths)
	}

	for _, v := range list {
		c.chartPakacge(v)
	}

}

func (c *Commands) chartPakacge(chartPath string) {

	var (
		chartOutput = viper.GetString("output.path")
		chartRepo   = viper.GetString("chart.repo")
		s3Force     = viper.GetBool("helm.s3.force")
		dryRun      = viper.GetBool("command.dry-run")
		separator1  = strings.Repeat("=", 80)
		separator2  = strings.Repeat("-", 80)
		argForce    = ""
		reportChart = Report{}
	)

	reportChart.ChartPath = chartPath
	m := helm.ChartVersion(chartPath)

	log.Infoln("|-> START <-|", separator2)

	reportChart.ChartName = m.Name
	reportChart.ChartVersion = m.Version
	reportChart.GitLsTree = viper.GetBool("git.lsTree")
	reportChart.Force = s3Force
	reportChart.RepoName = chartRepo
	reportChart.ChartURL = fmt.Sprintf("s3://%s/%s-%s.tgz", chartRepo, m.Name, m.Version)

	found, err := helm.Search(m.Name, m.Version)
	if err != nil {
		log.Fatalln(err)
	}

	if found && !s3Force {
		reportChart.Published = false
		reportPublish = append(reportPublish, reportChart)
		log.Warnf("|-> SKIPPING <-| The Helm Chart %s and %s version already exists!", m.Name, m.Version)
		log.Infoln("|->  END  <-|", separator1)
		return
	}

	reportChart.Published = true

	if err := helm.Package(m, chartPath, chartOutput); err != nil {
		reportChart.Published = false
		log.Fatalln(err)
	}

	if s3Force {
		reportChart.Force = true
		argForce = "--force"
	}

	log.Tracef("%#v, Force? %s", reportChart, argForce)

	if !dryRun {
		if err := plugins.S3Publisher(m, chartPath, chartRepo, chartOutput, argForce); err != nil {
			reportChart.Published = false
			log.Fatalln(err)
		}
	} else {
		log.Warnln("Dry run mode has been activated no publishing process will be executed!")
	}

	reportPublish = append(reportPublish, reportChart)

	log.Infoln("|->  END  <-|", separator1)
}

func (c *Commands) gitLsTree(paths []string) []string {

	log.Infoln("Git LS Tree is enabled!!!")

	var listPaths = git.MergeLsTree(paths)

	return listPaths
}

func (c *Commands) GenerateReport(report []Report) {

	var (
		reportType = viper.GetString("report.type")
		reportName = viper.GetString("report.name")
		reportPath = viper.GetString("report.path")
	)

	if reportType == "" {
		return
	}

	log.Infoln("|-> START :: REPORT <-| Creating the chart processing report...")

	switch reportType {
	case "json":
		reportName = fmt.Sprintf("%s/%s.json", reportPath, reportName)
		helpers.CreateFilePrettyJSON(report, reportName)
	case "text", "txt":
		reportName = fmt.Sprintf("%s/%s.txt", reportPath, reportName)
		helpers.CreateFileTextPlain(report, reportName)
	default:
		log.Errorf("Report extension [ %s ] not supported!", reportType)
	}

	log.Infof("|-> END :: REPORT <-| Report %s created successfully!!", reportName)

}
