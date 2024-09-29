package publish

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/toolsascode/helm-s3-publisher/internal/git"
	"github.com/toolsascode/helm-s3-publisher/internal/helm"
	"github.com/toolsascode/helm-s3-publisher/internal/plugins"
)

/*
	helm-s3-publisher REPO [CHART PATHS] [flags]
*/
// Run
func (c *Commands) Run() {

	var (
		listPath = viper.GetStringSlice("chart.paths")
	)

	c.chartList(listPath)

	// var chartPath = "/Users/carlosjunior/projects/winnin/helm-charts/prefect-worker"

	if err := helm.CheckIntall(); err != nil {
		log.Fatalln(err)
		// os.Exit(1)
	}
	if err := plugins.S3CheckIntall(); err != nil {
		log.Fatalln(err)
		// os.Exit(1)
	}

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
	)

	m := helm.ChartVersion(chartPath)

	if err := helm.Search(m.Name, m.Version); err != nil {
		log.Fatalln(err)
		// os.Exit(1)
	}
	if err := helm.Package(m, chartPath, chartOutput); err != nil {
		log.Fatalln(err)
		// os.Exit(1)
	}
	if err := plugins.S3Publisher(m, chartPath, chartRepo, chartOutput); err != nil {
		log.Fatalln(err)
		// os.Exit(1)
	}
}

func (c *Commands) gitLsTree(paths []string) []string {

	// if !viper.GetBool("git.lsTree") {
	// 	log.Infoln("Git LS Tree is not enabled!!!")
	// 	return []string{}
	// }

	log.Infoln("Git LS Tree is enabled!!!")

	var listPaths = git.MergeLsTree(paths)

	return listPaths
}
