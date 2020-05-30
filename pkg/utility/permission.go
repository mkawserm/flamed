package utility

import "github.com/mkawserm/flamed/pkg/pb"

func HasReadPermission(ac *pb.AccessControl) bool {
	if ac == nil {
		return false
	}

	if len(ac.Permission) == 0 {
		return false
	}

	var p = ac.Permission[0]

	return hasBit(p, 0)
}

func HasWritePermission(ac *pb.AccessControl) bool {
	if ac == nil {
		return false
	}

	if len(ac.Permission) == 0 {
		return false
	}

	var p = ac.Permission[0]

	return hasBit(p, 1)
}

func HasUpdatePermission(ac *pb.AccessControl) bool {
	if ac == nil {
		return false
	}

	if len(ac.Permission) == 0 {
		return false
	}

	var p = ac.Permission[0]

	return hasBit(p, 2)
}

func HasDeletePermission(ac *pb.AccessControl) bool {
	if ac == nil {
		return false
	}

	if len(ac.Permission) == 0 {
		return false
	}

	var p = ac.Permission[0]

	return hasBit(p, 3)
}

func HasGlobalSearchPermission(ac *pb.AccessControl) bool {
	if ac == nil {
		return false
	}

	if len(ac.Permission) == 0 {
		return false
	}

	var p = ac.Permission[0]

	return hasBit(p, 4)
}

func HasGlobalIteratePermission(ac *pb.AccessControl) bool {
	if ac == nil {
		return false
	}

	if len(ac.Permission) == 0 {
		return false
	}

	var p = ac.Permission[0]

	return hasBit(p, 5)
}

func HasGlobalRetrievePermission(ac *pb.AccessControl) bool {
	if ac == nil {
		return false
	}

	if len(ac.Permission) == 0 {
		return false
	}

	var p = ac.Permission[0]

	return hasBit(p, 6)
}

func NewPermission(read bool, write bool, update bool, delete bool,
	globalSearch bool,
	globalIterate bool,
	globalRetrieve bool) []byte {
	var p uint8 = 0

	if read {
		p = setBit(p, 0)
	}
	if write {
		p = setBit(p, 1)
	}
	if update {
		p = setBit(p, 2)
	}
	if delete {
		p = setBit(p, 3)
	}

	if globalSearch {
		p = setBit(p, 4)
	}

	if globalIterate {
		p = setBit(p, 5)
	}

	if globalRetrieve {
		p = setBit(p, 6)
	}

	return []byte{p}
}

func setBit(n uint8, pos uint8) uint8 {
	n |= 1 << pos
	return n
}

func hasBit(n uint8, pos uint8) bool {
	val := n & (1 << pos)
	return val > 0
}
