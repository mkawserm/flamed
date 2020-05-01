package utility

func IsNamespaceValid(namespace []byte) bool {
	if len(namespace) < 3 {
		return false
	}

	if namespace[0] >= 'A' && namespace[1] <= 'z' {
		return true
	}

	return false
}
