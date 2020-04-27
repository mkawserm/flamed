package uidutil

import (
	"bytes"
	"testing"
)

func TestGetUid(t *testing.T) {
	t.Helper()

	GetUid([]byte("1"), []byte("1"))
}

func TestGetUidString(t *testing.T) {
	t.Helper()

	GetUidString([]byte("1"), []byte("1"))
}

func TestGetNamespace(t *testing.T) {
	t.Helper()

	if !bytes.Equal(GetNamespace(GetUid([]byte("1"), []byte("2"))), []byte("1")) {
		t.Fatalf("namespace does not match")
	}
}

func TestGetUidFromString(t *testing.T) {
	t.Helper()

	GetUidFromString(GetUidString([]byte("1"), []byte("1")))
}

func TestGetNamespaceFromString(t *testing.T) {
	t.Helper()

	GetNamespaceFromString(GetUidString([]byte("1"), []byte("1")))
}

func TestSplitUid(t *testing.T) {
	t.Helper()

	n, k := SplitUid(GetUid([]byte("1"), []byte("2")))

	if !bytes.Equal(n, []byte("1")) {
		t.Fatalf("namespace does not match")
	}
	if !bytes.Equal(k, []byte("2")) {
		t.Fatalf("key does not match")
	}
}

func TestSplitUidString(t *testing.T) {
	t.Helper()

	n, k := SplitUidString(GetUidString([]byte("1"), []byte("2")))

	if !bytes.Equal(n, []byte("1")) {
		t.Fatalf("namespace does not match")
	}

	if !bytes.Equal(k, []byte("2")) {
		t.Fatalf("key does not match")
	}
}
