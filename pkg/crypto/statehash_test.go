package crypto

import "testing"

func TestGetStateHashFromStringKey(t *testing.T) {
	t.Helper()
	hash := GetStateHashFromStringKey("test", "1")
	if len(hash) != 35 {
		t.Errorf("unexpected state hash length, expected length 35, got %d", len(hash))
		t.Errorf("Hash: %v", hash)
	}
}

func TestGetStateHashFromUint64Key(t *testing.T) {
	t.Helper()
	hash := GetStateHashFromUint64Key("test", 1)
	if len(hash) != 35 {
		t.Errorf("unexpected state hash length, expected length 35, got %d", len(hash))
		t.Errorf("Hash: %v", hash)
	}
}
