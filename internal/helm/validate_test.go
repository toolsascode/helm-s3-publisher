package helm

import "testing"

func TestCheckIntall(t *testing.T) {
	if err := CheckIntall(); err != nil {
		t.Fatalf(`CheckInstall %v, want "", error`, err)
	}
}
