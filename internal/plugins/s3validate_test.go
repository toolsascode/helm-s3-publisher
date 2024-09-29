package plugins

import "testing"

func TestS3CheckIntall(t *testing.T) {
	if err := S3CheckIntall(); err != nil {
		t.Fatalf(`S3CheckIntall %v, want "", error`, err)
	}
}
