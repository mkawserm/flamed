package badger

import (
	"github.com/mkawserm/flamed/testsuite/storage"
	"testing"
)

func TestBadger(t *testing.T) {
	t.Helper()
	storage.StateStorageTestSuite(t, &Badger{})
}
