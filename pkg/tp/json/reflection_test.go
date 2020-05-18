package json

import "testing"

type myType struct {
	ID string `json:"id"`
}

func TestGetId(t *testing.T) {
	t.Helper()

	if GetId(&myType{ID: "1"}) != "1" {
		t.Error("unexpected id")
	}
}
