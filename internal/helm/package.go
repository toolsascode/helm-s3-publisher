package helm

import (
	"errors"
	"os/exec"

	"helm.sh/helm/v3/pkg/chart"

	log "github.com/sirupsen/logrus"
)

func Package(chart *chart.Metadata, chartPath, chartOutput string) error {

	var depFlag = ""

	if len(chart.Dependencies) > 0 {
		depFlag = "--dependency-update"
		log.Warn("The chart has dependencies will be updated in the packager.")
	}

	log.Infof("Packaging the chart %s and %s version", chart.Name, chart.Version)

	out, err := exec.Command("helm", "package", chartPath, "--destination", chartOutput, depFlag).Output()
	if errors.Is(err, exec.ErrDot) {
		log.Fatal(err)
		return err
	}

	log.Infoln("Release on chart processed successfully!!!")
	log.Infof("%s", out)

	return nil
}
