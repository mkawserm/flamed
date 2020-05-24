package utility

import "github.com/google/uuid"

func UUIDToString(ub []byte) string {
	u, e := uuid.FromBytes(ub)
	if e != nil {
		return ""
	}
	return u.String()
}
