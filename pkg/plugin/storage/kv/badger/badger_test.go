package badger

import (
	"github.com/dgraph-io/badger/v2"
	"github.com/mkawserm/flamed/pkg/x"
	"os"
	"testing"
)

func removeAll(path string) {
	_ = os.RemoveAll(path)
}

func TestBadger_Open(t *testing.T) {
	t.Helper()

	path := "/tmp/badger_test"
	secretKey := []byte("")

	defer removeAll(path)

	b := &Badger{}
	o, err := b.Open(path, secretKey, false, nil)

	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}

	if !o {
		t.Fatalf("failed to open database")
	}

	// Second time open to execute first branch
	o, err = b.Open(path, secretKey, false, nil)

	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}

	if !o {
		t.Fatalf("failed to open database")
	}

	err = b.Close()

	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}
}

func TestBadger_OpenWithConfiguration(t *testing.T) {
	t.Helper()

	path := "/tmp/badger_test"
	secretKey := []byte("")

	defer removeAll(path)

	b := &Badger{}
	conf := Configuration{
		SliceCap:                      0,
		LogPrefix:                     "",
		GoroutineNumber:               0,
		BadgerOptions:                 badger.DefaultOptions(""),
		EncryptionKeyRotationDuration: 0,
	}

	o, err := b.Open(path, secretKey, false, conf)

	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}

	if !o {
		t.Fatalf("failed to open database")
	}

	err = b.Close()

	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}
}

func TestBadger_OpenFailure(t *testing.T) {
	t.Helper()

	path := "/tmp/badger_test"
	secretKey := []byte("")

	defer removeAll(path)

	b := &Badger{}
	conf := Configuration{
		SliceCap:                      0,
		LogPrefix:                     "",
		GoroutineNumber:               0,
		BadgerOptions:                 badger.DefaultOptions("").WithValueThreshold((1 << 20) + 1),
		EncryptionKeyRotationDuration: 0,
	}

	_, err := b.Open(path, secretKey, false, conf)

	if err != x.ErrFailedToOpenStorage {
		t.Fatalf("Unexpected error %v", err)
	}

	err = b.Close()

	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}

}
