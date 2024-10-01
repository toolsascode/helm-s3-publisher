package plugins

import (
	"testing"

	"helm.sh/helm/v3/pkg/chart"
)

func TestS3Publisher(t *testing.T) {

	var chart = &chart.Metadata{
		Name:    "argocd",
		Version: "7.15.0",
	}

	if err := S3Publisher(chart, "./", "local", "./"); err != nil {
		t.Fatalf(`S3Publisher %v, want "", error`, err)
	}
}
