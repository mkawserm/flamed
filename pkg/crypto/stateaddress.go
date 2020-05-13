package crypto

import (
	"crypto/sha512"
)

// GetStateAddressFromStringKey generate 35 bytes length state address from string
// based key
func GetStateAddressFromStringKey(familyName, key string) []byte {
	familyNameHash := sha512.Sum512([]byte(familyName))
	keyHash := sha512.Sum512([]byte(key))
	r := make([]byte, 35)
	r = familyNameHash[0:3]
	copy(r[3:], keyHash[32:])
	return r
}
