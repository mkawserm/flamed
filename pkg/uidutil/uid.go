package uidutil

import "bytes"
import "encoding/hex"

func GetUID(namespace []byte, key []byte) []byte {
	if len(key) == 0 {
		return namespace[:]
	} else {
		r := make([]byte, 0, len(namespace)+1+len(key))
		r = append(r, namespace...)
		r = append(r, ':')
		r = append(r, key...)
		return r
	}
}

func GetUIDString(namespace []byte, key []byte) string {
	src := GetUID(namespace, key)
	return hex.EncodeToString(src)
}

func GetUIDFromString(uidString string) []byte {
	r, _ := hex.DecodeString(uidString)
	return r
}

func GetNamespace(uid []byte) []byte {
	r := bytes.Split(uid, []byte(":"))
	if len(r) == 0 {
		return nil
	} else {
		return r[0]
	}
}

func GetNamespaceFromString(uidString string) []byte {
	return GetNamespace(GetUIDFromString(uidString))
}
