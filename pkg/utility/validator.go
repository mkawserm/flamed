package utility

import "bytes"

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
