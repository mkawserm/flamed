package utility

import "testing"

func TestIsNamespaceValid(t *testing.T) {
	t.Helper()

	if IsNamespaceValid([]byte("1")) {
		t.Fatalf("Namespace should be invalid for `1`")
	}

	if IsNamespaceValid([]byte("abc::xyz")) {
		t.Fatalf("Namespace should be invalid for `abc::xyz`")
	}

	if IsNamespaceValid([]byte("123")) {
		t.Fatalf("Namespace should be invalid for `123`")
	}

	if !IsNamespaceValid([]byte("abc")) {
		t.Fatalf("Namespace should be valid for `abc`")
	}

	if !IsNamespaceValid([]byte("Abc")) {
		t.Fatalf("Namespace should be valid for `Abc`")
	}

	if !IsNamespaceValid([]byte("SMILE")) {
		t.Fatalf("Namespace should be valid for `SMILE`")
	}

	if !IsNamespaceValid([]byte("yellow")) {
		t.Fatalf("Namespace should be valid for `yellow`")
	}
}

func TestIsUsernameValid(t *testing.T) {
	t.Helper()

	if IsUsernameValid("t") {
		t.Fatalf("usernmae should be invalid for `t`")
	}

	if IsUsernameValid("test::user") {
		t.Fatalf("usernmae should be invalid for `test::user`")
	}

	if !IsUsernameValid("test_user") {
		t.Fatalf("usernmae should be invalid for `test_user`")
	}
}

func TestIsPasswordValid(t *testing.T) {
	t.Helper()

	if IsPasswordValid("123") {
		t.Fatalf("passwod should be invalid for `123`")
	}

	if !IsPasswordValid("123456") {
		t.Fatalf("passwod should be invalid for `123456`")
	}
}
