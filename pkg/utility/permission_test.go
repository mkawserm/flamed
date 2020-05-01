package utility

import (
	"github.com/mkawserm/flamed/pkg/pb"
	"testing"
)

func TestHasReadPermission(t *testing.T) {
	t.Helper()

	p := NewPermission(true, false, false, false)
	ac := &pb.FlameAccessControl{
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
}

func TestHasWritePermission(t *testing.T) {
	t.Helper()

	p := NewPermission(false, true, false, false)
	ac := &pb.FlameAccessControl{
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
}

func TestHasUpdatePermission(t *testing.T) {
	t.Helper()

	p := NewPermission(false, false, true, false)
	ac := &pb.FlameAccessControl{
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
}

func TestHasDeletePermission(t *testing.T) {
	t.Helper()

	p := NewPermission(false, false, false, true)
	ac := &pb.FlameAccessControl{
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
}

func TestNewPermission(t *testing.T) {
	t.Helper()

	p := NewPermission(false, false, false, false)
	ac := &pb.FlameAccessControl{
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
}

func TestNewPermission2(t *testing.T) {
	t.Helper()

	p := NewPermission(true, true, true, true)
	ac := &pb.FlameAccessControl{
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
}
