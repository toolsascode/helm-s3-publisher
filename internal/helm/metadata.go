package helm

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"helm.sh/helm/v3/pkg/chart"
	"sigs.k8s.io/yaml"
)

func ChartVersion(chartPath string) *chart.Metadata {

	var chartMetadata = new(chart.Metadata)

	chartFile, err := os.ReadFile(fmt.Sprintf("%s/%s", chartPath, "Chart.yaml"))
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
		return &chart.Metadata{}
	}

	if err := yaml.Unmarshal(chartFile, chartMetadata); err != nil {
		log.Fatal("Cannot load Chart.yaml.")
		return &chart.Metadata{}
	}

	return chartMetadata

}
