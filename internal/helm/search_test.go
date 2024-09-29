package helm

import "testing"

func TestSearch(t *testing.T) {
	if err := Search("test", "1.0.0"); err != nil {
		t.Fatalf(`Search %v, want "", error`, err)
	}
}
