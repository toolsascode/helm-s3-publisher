package helm

import (
	"testing"

	"helm.sh/helm/v3/pkg/chart"
)

func TestPackage(t *testing.T) {

	var chart = &chart.Metadata{
		Name:    "test",
		Version: "1.0.0",
	}

	if err := Package(chart, "./", "./"); err != nil {
		t.Fatalf(`Search %v, want "", error`, err)
	}
}
