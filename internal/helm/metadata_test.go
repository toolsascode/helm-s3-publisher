package helm

import (
	"testing"
)

func TestMetadata(t *testing.T) {
	if chart := ChartVersion("."); chart == nil {
		t.Fatalf(`Search %v, want "", error`, "Empty")
	}
}
