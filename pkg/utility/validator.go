package utility

import (
	"bytes"
	"strings"
)

func IsNamespaceValid(namespace []byte) bool {
	if len(namespace) < 3 {
		return false
	}

	if bytes.Contains(namespace, []byte("::")) {
		return false
	}

	if namespace[0] >= 'A' && namespace[1] <= 'z' {
		return true
	}

	return false
}

func IsUsernameValid(username string) bool {
	if strings.Contains(username, "::") {
		return false
	}
	return len(username) >= 3
}

func IsPasswordValid(password string) bool {
	return len(password) >= 6
}
