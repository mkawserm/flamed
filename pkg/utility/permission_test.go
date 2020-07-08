package utility

import (
	"github.com/mkawserm/flamed/pkg/pb"
	"testing"
)

func TestHasReadPermission(t *testing.T) {
	t.Helper()

	p := NewPermission(true,
		false,
		false,
		false,
		false,
		false,
		false,
		false)
	ac := &pb.AccessControl{
		Permission: p,
	}

	if !HasReadPermission(ac) {
		t.Fatalf("Unexpected read permission")
	}

	if HasWritePermission(ac) {
		t.Fatalf("Unexpected write permission")
	}

	if HasUpdatePermission(ac) {
		t.Fatalf("Unexpected update permission")
	}

	if HasDeletePermission(ac) {
		t.Fatalf("Unexpected delete permission")
	}

	if HasGlobalSearchPermission(ac) {
		t.Fatalf("Unexpected global search permission")
	}

	if HasGlobalIteratePermission(ac) {
		t.Fatalf("Unexpected global iterate permission")
	}

	if HasGlobalRetrievePermission(ac) {
		t.Fatalf("Unexpected global retreive permission")
	}

	if HasGlobalCRUDPermission(ac) {
		t.Fatalf("Unexpected global crud permission")
	}
}

func TestHasWritePermission(t *testing.T) {
	t.Helper()

	p := NewPermission(false,
		true,
		false,
		false,
		false,
		false,
		false,
		false)
	ac := &pb.AccessControl{
		Permission: p,
	}

	if HasReadPermission(ac) {
		t.Fatalf("Unexpected read permission")
	}

	if !HasWritePermission(ac) {
		t.Fatalf("Unexpected write permission")
	}

	if HasUpdatePermission(ac) {
		t.Fatalf("Unexpected update permission")
	}

	if HasDeletePermission(ac) {
		t.Fatalf("Unexpected delete permission")
	}

	if HasGlobalSearchPermission(ac) {
		t.Fatalf("Unexpected global search permission")
	}

	if HasGlobalIteratePermission(ac) {
		t.Fatalf("Unexpected global iterate permission")
	}

	if HasGlobalRetrievePermission(ac) {
		t.Fatalf("Unexpected global retreive permission")
	}

	if HasGlobalCRUDPermission(ac) {
		t.Fatalf("Unexpected global crud permission")
	}
}

func TestHasUpdatePermission(t *testing.T) {
	t.Helper()

	p := NewPermission(false,
		false,
		true,
		false,
		false,
		false,
		false,
		false)
	ac := &pb.AccessControl{
		Permission: p,
	}

	if HasReadPermission(ac) {
		t.Fatalf("Unexpected read permission")
	}

	if HasWritePermission(ac) {
		t.Fatalf("Unexpected write permission")
	}

	if !HasUpdatePermission(ac) {
		t.Fatalf("Unexpected update permission")
	}

	if HasDeletePermission(ac) {
		t.Fatalf("Unexpected delete permission")
	}

	if HasGlobalSearchPermission(ac) {
		t.Fatalf("Unexpected global search permission")
	}

	if HasGlobalIteratePermission(ac) {
		t.Fatalf("Unexpected global iterate permission")
	}

	if HasGlobalRetrievePermission(ac) {
		t.Fatalf("Unexpected global retreive permission")
	}

	if HasGlobalCRUDPermission(ac) {
		t.Fatalf("Unexpected global crud permission")
	}
}

func TestHasDeletePermission(t *testing.T) {
	t.Helper()

	p := NewPermission(false,
		false,
		false,
		true,
		false,
		false,
		false,
		false)
	ac := &pb.AccessControl{
		Permission: p,
	}

	if HasReadPermission(ac) {
		t.Fatalf("Unexpected read permission")
	}

	if HasWritePermission(ac) {
		t.Fatalf("Unexpected write permission")
	}

	if HasUpdatePermission(ac) {
		t.Fatalf("Unexpected update permission")
	}

	if !HasDeletePermission(ac) {
		t.Fatalf("Unexpected delete permission")
	}

	if HasGlobalSearchPermission(ac) {
		t.Fatalf("Unexpected global search permission")
	}

	if HasGlobalIteratePermission(ac) {
		t.Fatalf("Unexpected global iterate permission")
	}

	if HasGlobalRetrievePermission(ac) {
		t.Fatalf("Unexpected global retreive permission")
	}

	if HasGlobalCRUDPermission(ac) {
		t.Fatalf("Unexpected global crud permission")
	}
}

func TestHasGlobalSearchPermission(t *testing.T) {
	t.Helper()

	p := NewPermission(false,
		false,
		false,
		false,
		true,
		false,
		false,
		false)
	ac := &pb.AccessControl{
		Permission: p,
	}

	if HasReadPermission(ac) {
		t.Fatalf("Unexpected read permission")
	}

	if HasWritePermission(ac) {
		t.Fatalf("Unexpected write permission")
	}

	if HasUpdatePermission(ac) {
		t.Fatalf("Unexpected update permission")
	}

	if HasDeletePermission(ac) {
		t.Fatalf("Unexpected delete permission")
	}

	if !HasGlobalSearchPermission(ac) {
		t.Fatalf("Unexpected global search permission")
	}

	if HasGlobalIteratePermission(ac) {
		t.Fatalf("Unexpected global iterate permission")
	}

	if HasGlobalRetrievePermission(ac) {
		t.Fatalf("Unexpected global retreive permission")
	}

	if HasGlobalCRUDPermission(ac) {
		t.Fatalf("Unexpected global crud permission")
	}
}

func TestHasGlobalIteratePermission(t *testing.T) {
	t.Helper()

	p := NewPermission(false,
		false,
		false,
		false,
		false,
		true,
		false,
		false)
	ac := &pb.AccessControl{
		Permission: p,
	}

	if HasReadPermission(ac) {
		t.Fatalf("Unexpected read permission")
	}

	if HasWritePermission(ac) {
		t.Fatalf("Unexpected write permission")
	}

	if HasUpdatePermission(ac) {
		t.Fatalf("Unexpected update permission")
	}

	if HasDeletePermission(ac) {
		t.Fatalf("Unexpected delete permission")
	}

	if HasGlobalSearchPermission(ac) {
		t.Fatalf("Unexpected global search permission")
	}

	if !HasGlobalIteratePermission(ac) {
		t.Fatalf("Unexpected global iterate permission")
	}

	if HasGlobalRetrievePermission(ac) {
		t.Fatalf("Unexpected global retreive permission")
	}

	if HasGlobalCRUDPermission(ac) {
		t.Fatalf("Unexpected global crud permission")
	}
}

func TestHasGlobalRetrievePermission(t *testing.T) {
	t.Helper()

	p := NewPermission(false,
		false,
		false,
		false,
		false,
		false,
		true,
		false)
	ac := &pb.AccessControl{
		Permission: p,
	}

	if HasReadPermission(ac) {
		t.Fatalf("Unexpected read permission")
	}

	if HasWritePermission(ac) {
		t.Fatalf("Unexpected write permission")
	}

	if HasUpdatePermission(ac) {
		t.Fatalf("Unexpected update permission")
	}

	if HasDeletePermission(ac) {
		t.Fatalf("Unexpected delete permission")
	}

	if HasGlobalSearchPermission(ac) {
		t.Fatalf("Unexpected global search permission")
	}

	if HasGlobalIteratePermission(ac) {
		t.Fatalf("Unexpected global iterate permission")
	}

	if !HasGlobalRetrievePermission(ac) {
		t.Fatalf("Unexpected global retreive permission")
	}

	if HasGlobalCRUDPermission(ac) {
		t.Fatalf("Unexpected global crud permission")
	}
}

func TestHasGlobalCRUDPermission(t *testing.T) {
	t.Helper()

	p := NewPermission(false,
		false,
		false,
		false,
		false,
		false,
		false,
		true)
	ac := &pb.AccessControl{
		Permission: p,
	}

	if HasReadPermission(ac) {
		t.Fatalf("Unexpected read permission")
	}

	if HasWritePermission(ac) {
		t.Fatalf("Unexpected write permission")
	}

	if HasUpdatePermission(ac) {
		t.Fatalf("Unexpected update permission")
	}

	if HasDeletePermission(ac) {
		t.Fatalf("Unexpected delete permission")
	}

	if HasGlobalSearchPermission(ac) {
		t.Fatalf("Unexpected global search permission")
	}

	if HasGlobalIteratePermission(ac) {
		t.Fatalf("Unexpected global iterate permission")
	}

	if HasGlobalRetrievePermission(ac) {
		t.Fatalf("Unexpected global retreive permission")
	}

	if !HasGlobalCRUDPermission(ac) {
		t.Fatalf("Unexpected global crud permission")
	}
}

func TestNewPermission(t *testing.T) {
	t.Helper()

	p := NewPermission(false,
		false,
		false,
		false,
		false,
		false,
		false,
		false)
	ac := &pb.AccessControl{
		Permission: p,
	}

	if HasReadPermission(ac) {
		t.Fatalf("Unexpected read permission")
	}

	if HasWritePermission(ac) {
		t.Fatalf("Unexpected write permission")
	}

	if HasUpdatePermission(ac) {
		t.Fatalf("Unexpected update permission")
	}

	if HasDeletePermission(ac) {
		t.Fatalf("Unexpected delete permission")
	}

	if HasGlobalSearchPermission(ac) {
		t.Fatalf("Unexpected global search permission")
	}

	if HasGlobalIteratePermission(ac) {
		t.Fatalf("Unexpected global iterate permission")
	}

	if HasGlobalRetrievePermission(ac) {
		t.Fatalf("Unexpected global retreive permission")
	}

	if HasGlobalCRUDPermission(ac) {
		t.Fatalf("Unexpected global crud permission")
	}
}

func TestNewPermission2(t *testing.T) {
	t.Helper()

	p := NewPermission(true,
		true,
		true,
		true,
		true,
		true,
		true,
		true)
	ac := &pb.AccessControl{
		Permission: p,
	}

	if !HasReadPermission(ac) {
		t.Fatalf("Unexpected read permission")
	}

	if !HasWritePermission(ac) {
		t.Fatalf("Unexpected write permission")
	}

	if !HasUpdatePermission(ac) {
		t.Fatalf("Unexpected update permission")
	}

	if !HasDeletePermission(ac) {
		t.Fatalf("Unexpected delete permission")
	}

	if !HasGlobalSearchPermission(ac) {
		t.Fatalf("Unexpected global search permission")
	}

	if !HasGlobalIteratePermission(ac) {
		t.Fatalf("Unexpected global iterate permission")
	}

	if !HasGlobalRetrievePermission(ac) {
		t.Fatalf("Unexpected global retreive permission")
	}

	if !HasGlobalCRUDPermission(ac) {
		t.Fatalf("Unexpected global crud permission")
	}
}
