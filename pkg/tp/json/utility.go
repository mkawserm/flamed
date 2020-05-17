package json

import "github.com/mkawserm/flamed/pkg/crypto"

const separator = "::"

func GetJSONFamilyStateAddress(namespace []byte, family, id string) []byte {
	return crypto.GetStateAddress(namespace, []byte(family+separator+id))
}
