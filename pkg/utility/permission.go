package utility

import "github.com/mkawserm/flamed/pkg/pb"

func HasReadPermission(ac *pb.FlameAccessControl) bool {
	if ac == nil {
		return false
	}

	if len(ac.Permission) == 0 {
		return false
	}

	var p = ac.Permission[0]

	return hasBit(p, 0)
}

func HasWritePermission(ac *pb.FlameAccessControl) bool {
	if ac == nil {
		return false
	}

	if len(ac.Permission) == 0 {
		return false
	}

	var p = ac.Permission[0]

	return hasBit(p, 1)
}

func HasUpdatePermission(ac *pb.FlameAccessControl) bool {
	if ac == nil {
		return false
	}

	if len(ac.Permission) == 0 {
		return false
	}

	var p = ac.Permission[0]

	return hasBit(p, 2)
}

func HasDeletePermission(ac *pb.FlameAccessControl) bool {
	if ac == nil {
		return false
	}

	if len(ac.Permission) == 0 {
		return false
	}

	var p = ac.Permission[0]

	return hasBit(p, 3)
}

func NewPermission(read bool, write bool, update bool, delete bool) []byte {
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

	return []byte{p}
}

func setBit(n uint8, pos uint8) uint8 {
	n |= 1 << pos
	return n
}

//func clearBit(n uint8, pos uint8) uint8 {
//	var mask uint8 = ^(1 << pos)
//	n &= mask
//	return n
//}

func hasBit(n uint8, pos uint8) bool {
	val := n & (1 << pos)
	return val > 0
}
