package json

import "github.com/mkawserm/flamed/pkg/crypto"

func GetJSONFamilyStateAddress(namespace []byte, family, id string) []byte {
	return crypto.GetStateAddress(namespace, []byte(family+"::"+id))
}
