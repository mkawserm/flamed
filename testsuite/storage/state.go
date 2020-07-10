package storage

import (
	"bytes"
	"github.com/mkawserm/flamed/pkg/iface"
	"os"
	"testing"
)

func StateStorageTestSuite(t *testing.T, stateStorage iface.IStateStorage) {
	path := "/tmp/test_db_1"
	defer func() {
		_ = os.RemoveAll(path)
	}()

	stateStorage.Setup(path, nil, nil)
	if err := stateStorage.Open(); err != nil {
		t.Fatal("unexpected error: ", err)
		return
	}

	defer func() {
		_ = stateStorage.Close()
	}()

	txn := stateStorage.NewTransaction()

	if txn == nil {
		t.Fatal("unexpected nil pointer")
		return
	}

	inputDataTable := []string{
		"z",
		"a",
		"Z",
		"A",
		"9",
		"0",
		"5",
		"1",
		"Ab",
		"1ba",
		"1ab",
	}

	expectedForwardDataTable := []string{
		"0",
		"1",
		"1ab",
		"1ba",
		"5",
		"9",
		"A",
		"Ab",
		"Z",
		"a",
		"z",
	}

	expectedReverseDataTable := []string{
		"z",
		"a",
		"Z",
		"Ab",
		"A",
		"9",
		"5",
		"1ba",
		"1ab",
		"1",
		"0",
	}

	for _, v := range inputDataTable {
		if err := txn.Set([]byte(v), []byte(v)); err != nil {
			t.Fatal("unexpected error: ", err)
			return
		}
	}

	if err := txn.Commit(); err != nil {
		t.Fatal("unexpected error: ", err)
		return
	}

	txn = stateStorage.NewTransaction()
	forwardIterator := txn.ForwardIterator()
	if forwardIterator == nil {
		t.Fatal("unexpected nil forward iterator")
		return
	}
	var i = 0
	for forwardIterator.Rewind(); forwardIterator.Valid(); forwardIterator.Next() {
		state := forwardIterator.StateSnapshot()
		currentData := expectedForwardDataTable[i]

		if !bytes.EqualFold([]byte(currentData), state.Address) {
			t.Fatal("address ordering is not correct")
		}

		if !bytes.EqualFold([]byte(currentData), state.Data) {
			t.Fatal("data mismatch")
		}

		i = i + 1
	}

	forwardIterator.Close()

	reverseIterator := txn.ReverseIterator()
	if reverseIterator == nil {
		t.Error("unexpected nil reverse iterator")
		return
	}
	i = 0
	for reverseIterator.Rewind(); reverseIterator.Valid(); reverseIterator.Next() {
		state := reverseIterator.StateSnapshot()
		currentData := expectedReverseDataTable[i]

		if !bytes.EqualFold([]byte(currentData), state.Address) {
			t.Fatal("address ordering is not correct")
		}

		if !bytes.EqualFold([]byte(currentData), state.Data) {
			t.Fatal("data mismatch")
		}

		i = i + 1
	}

	reverseIterator.Close()
}
