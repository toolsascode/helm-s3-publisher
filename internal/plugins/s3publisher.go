package plugins

import (
	"errors"
	"fmt"
	"os/exec"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/toolsascode/helm-s3-publisher/internal/helpers"
	"helm.sh/helm/v3/pkg/chart"
)

func S3Publisher(chart *chart.Metadata, chartPath, chartRepo, chartOutput string, args ...string) error {

	dryRun := viper.GetBool("command.dry-run")
	listArgs := helpers.MergeArgs(
		[]string{
			"s3",
			"push",
			fmt.Sprintf("%s/%s-%s.tgz", chartOutput, chart.Name, chart.Version),
			chartRepo,
		},
		args...)

	if len(viper.GetString("helm.s3.content-type")) > 0 {
		extraArg := []string{
			"--content-type",
			viper.GetString("helm.s3.content-type"),
		}
		listArgs = helpers.MergeArgs(listArgs, extraArg...)
	}

	if len(viper.GetString("helm.s3.acl")) > 0 {
		extraArg := []string{
			"--acl",
			viper.GetString("helm.s3.acl"),
		}
		listArgs = helpers.MergeArgs(listArgs, extraArg...)
	}

	log.Tracef("helpers.MergeArgs: %#v", listArgs)

	log.Infof("The chart publishing process has started:\nName: %s\nVersion: %s\nRepository: %s\nLocated at: %s", chart.Name, chart.Version, chartRepo, chartPath)

	if dryRun {
		log.Warnln("Dry run mode has been activated no publishing process will be executed!")
		return nil
	}

	out, err := exec.Command("helm", listArgs...).Output()
	if errors.Is(err, exec.ErrDot) {
		log.Fatal(err, out)
		return err
	}

	if err != nil {
		log.Errorf("S3Publisher :: Unable to publish version %s of package %s check if the specified repository %s exists and has permission to publish.", chart.Version, chart.Name, chartRepo)
		log.Fatalf("S3Publisher - Err: %v :Out: %v", err, out)
		return err
	}

	log.Infoln("The chart was published successfully!!!")
	log.Infof("%s", out)

	return nil
}
