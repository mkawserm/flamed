package badger

import (
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
