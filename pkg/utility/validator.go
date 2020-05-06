package utility

import (
	"bytes"
	"github.com/mkawserm/flamed/pkg/pb"
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
	return len(username) >= 3
}

func IsFlameUserValid(user *pb.FlameUser) bool {
	if len(user.Username) >= 3 && len(user.Password) >= 6 {
		return true
	} else {
		return false
	}
}
