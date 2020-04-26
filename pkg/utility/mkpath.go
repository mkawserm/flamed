package utility

import (
	"os"
)

// MkPath creates the path if not exists
func MkPath(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err = os.MkdirAll(path, os.FileMode(0700))
		if err == nil {
			return true
		} else {
			return false
		}
	}
	return true
}
