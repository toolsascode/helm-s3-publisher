package helm

import "testing"

func TestSearch(t *testing.T) {
	if err := Search("argo-cd", "7.15.0"); err != nil {
		t.Fatalf(`Search %v, want "", error`, err)
	}
}
