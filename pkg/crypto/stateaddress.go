package crypto

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"github.com/mkawserm/flamed/pkg/constant"
)

// GetStateAddress generates state address
// from namespace and key
func GetStateAddress(namespace []byte, key []byte) []byte {
	if len(namespace) == 0 {
		return nil
	} else if len(key) == 0 {
		return namespace[:]
	} else {
		r := make([]byte, 0, len(namespace)+1+len(key))
		r = append(r, namespace...)
		r = append(r, constant.Separator...)
		r = append(r, key...)
		return r
	}
}

// GetStateAddressHexString generates hex state address
// from namespace and key
func GetStateAddressHexString(namespace []byte, key []byte) string {
	src := GetStateAddress(namespace, key)
	return hex.EncodeToString(src)
}

// StateAddressByteSliceToHexString converts
// byte slice to hex string
func StateAddressByteSliceToHexString(address []byte) string {
	return hex.EncodeToString(address)
}

// StateAddressHexStringToByteSlice converts
// hex address string to byte slice
func StateAddressHexStringToByteSlice(address string) []byte {
	b, err := hex.DecodeString(address)
	if err != nil {
		return nil
	}
	return b
}

// GetStateAddressFromHexString converts state hex
// string to byte slice
func GetStateAddressFromHexString(stateAddressHexString string) []byte {
	r, _ := hex.DecodeString(stateAddressHexString)
	return r
}

// GetNamespace from state address byte slice
func GetNamespace(stateAddressByteSlice []byte) []byte {
	r := bytes.Split(stateAddressByteSlice, []byte(constant.Separator))
	if len(r) == 0 {
		return nil
	} else {
		return r[0]
	}
}

// GetNamespaceFromStateAddressHexString returns address byte slice
// from state address hex string
func GetNamespaceFromStateAddressHexString(stateAddressHexString string) []byte {
	return GetNamespace(GetStateAddressFromHexString(stateAddressHexString))
}

// SplitStateAddress splits state address to
// namespace and key
func SplitStateAddress(address []byte) ([]byte, []byte) {
	r := bytes.Split(address, []byte(constant.Separator))
	if len(r) == 0 {
		return nil, nil
	} else if len(r) == 1 {
		return r[0], nil
	} else if len(r) == 2 {
		return r[0], r[1]
	} else {
		key := bytes.TrimPrefix(address, r[0])
		key = bytes.TrimPrefix(key, []byte(constant.Separator))
		return r[0], key
	}
}

// SplitStateAddressHexString splits state address hex string
// to namespace and key
func SplitStateAddressHexString(stateAddressHexString string) ([]byte, []byte) {
	return SplitStateAddress(GetStateAddressFromHexString(stateAddressHexString))
}

// Uint64ToByteSlice converts uint64 to byte slice
func Uint64ToByteSlice(u uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b[0:], u)
	return b
}

// ByteSliceToUint64 converts byte slice to uint64
func ByteSliceToUint64(b []byte) uint64 {
	n := binary.BigEndian.Uint64(b[0:])
	return n
}
