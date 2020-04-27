package uidutil

import (
	"bytes"
	"testing"
)

func TestGetUID(t *testing.T) {
	t.Helper()

	GetUID([]byte("1"), []byte("1"))
}

func TestGetUIDString(t *testing.T) {
	t.Helper()

	GetUIDString([]byte("1"), []byte("1"))
}

func TestGetNamespace(t *testing.T) {
	t.Helper()

	if !bytes.Equal(GetNamespace(GetUID([]byte("1"), []byte("2"))), []byte("1")) {
		t.Fatalf("namespace does not match")
	}
}

func TestGetUIDFromString(t *testing.T) {
	t.Helper()

	GetUIDFromString(GetUIDString([]byte("1"), []byte("1")))
}

func TestGetNamespaceFromString(t *testing.T) {
	t.Helper()

	GetNamespaceFromString(GetUIDString([]byte("1"), []byte("1")))
}
