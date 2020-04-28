package uidutil

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
)

const separator string = "::"

func GetUid(namespace []byte, key []byte) []byte {
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

func GetUidString(namespace []byte, key []byte) string {
	src := GetUid(namespace, key)
	return hex.EncodeToString(src)
}

func GetUidFromString(uidString string) []byte {
	r, _ := hex.DecodeString(uidString)
	return r
}

func GetNamespace(uid []byte) []byte {
	r := bytes.Split(uid, []byte(separator))
	if len(r) == 0 {
		return nil
	} else {
		return r[0]
	}
}

func GetNamespaceFromString(uidString string) []byte {
	return GetNamespace(GetUidFromString(uidString))
}

func SplitUid(uid []byte) ([]byte, []byte) {
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

func SplitUidString(uidString string) ([]byte, []byte) {
	return SplitUid(GetUidFromString(uidString))
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
