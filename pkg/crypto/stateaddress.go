package crypto

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
)

const separator string = "::"

func GetStateAddress(namespace []byte, key []byte) []byte {
	if len(namespace) == 0 {
		return nil
	} else if len(key) == 0 {
		return namespace[:]
	} else {
		r := make([]byte, 0, len(namespace)+1+len(key))
		r = append(r, namespace...)
		r = append(r, separator...)
		r = append(r, key...)
		return r
	}
}

func GetStateAddressHexString(namespace []byte, key []byte) string {
	src := GetStateAddress(namespace, key)
	return hex.EncodeToString(src)
}

func StateAddressByteSliceToHexString(address []byte) string {
	return hex.EncodeToString(address)
}

func StateAddressHexStringToByteSlice(address string) []byte {
	b, err := hex.DecodeString(address)
	if err != nil {
		return nil
	}
	return b
}

func GetStateAddressFromHexString(stateAddressHexString string) []byte {
	r, _ := hex.DecodeString(stateAddressHexString)
	return r
}

func GetNamespace(stateAddressByteSlice []byte) []byte {
	r := bytes.Split(stateAddressByteSlice, []byte(separator))
	if len(r) == 0 {
		return nil
	} else {
		return r[0]
	}
}

func GetNamespaceFromStateAddressHexString(stateAddressHexString string) []byte {
	return GetNamespace(GetStateAddressFromHexString(stateAddressHexString))
}

func SplitStateAddress(uid []byte) ([]byte, []byte) {
	r := bytes.Split(uid, []byte(separator))
	if len(r) == 0 {
		return nil, nil
	} else if len(r) == 1 {
		return r[0], nil
	} else if len(r) == 2 {
		return r[0], r[1]
	} else {
		key := bytes.TrimPrefix(uid, r[0])
		key = bytes.TrimPrefix(key, []byte(separator))
		return r[0], key
	}
}

func SplitStateAddressHexString(stateAddressHexString string) ([]byte, []byte) {
	return SplitStateAddress(GetStateAddressFromHexString(stateAddressHexString))
}

func Uint64ToByteSlice(u uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b[0:], u)

	return b
}

func ByteSliceToUint64(b []byte) uint64 {
	n := binary.BigEndian.Uint64(b[0:])
	return n
}
