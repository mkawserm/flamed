package utility

import (
	"os"
	"testing"
)

func TestMkPath(t *testing.T) {
	t.Helper()

	if !MkPath("/tmp/1") {
		t.Fatalf("`%s` does not exists", "/tmp/1")
	} else {
		_ = os.RemoveAll("/tmp/1")
	}
}
