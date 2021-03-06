package utility

import "github.com/mkawserm/flamed/pkg/pb"

// HasReadPermission checks availability of read permission
func HasReadPermission(ac *pb.AccessControl) bool {
	if ac == nil {
		return false
	}
	return hasBit(ac.Permission, 0)
}

// HasWritePermission checks availability of write permission
func HasWritePermission(ac *pb.AccessControl) bool {
	if ac == nil {
		return false
	}
	return hasBit(ac.Permission, 1)
}

// HasUpdatePermission checks availability of update permission
func HasUpdatePermission(ac *pb.AccessControl) bool {
	if ac == nil {
		return false
	}
	return hasBit(ac.Permission, 2)
}

// HasDeletePermission checks availability of delete permission
func HasDeletePermission(ac *pb.AccessControl) bool {
	if ac == nil {
		return false
	}
	return hasBit(ac.Permission, 3)
}

// HasGlobalSearchPermission checks availability of global search permission
func HasGlobalSearchPermission(ac *pb.AccessControl) bool {
	if ac == nil {
		return false
	}
	return hasBit(ac.Permission, 4)
}

// HasGlobalIteratePermission checks availability of global iterate permission
func HasGlobalIteratePermission(ac *pb.AccessControl) bool {
	if ac == nil {
		return false
	}
	return hasBit(ac.Permission, 5)
}

// HasGlobalRetrievePermission checks availability of global retrieve permission
func HasGlobalRetrievePermission(ac *pb.AccessControl) bool {
	if ac == nil {
		return false
	}
	return hasBit(ac.Permission, 6)
}

// HasGlobalCRUDPermission checks availability of global CRUD permission
func HasGlobalCRUDPermission(ac *pb.AccessControl) bool {
	if ac == nil {
		return false
	}
	return hasBit(ac.Permission, 7)
}

// NewPermission creates permission unsigned integer based on
// parameters boolean flag
func NewPermission(read bool, write bool, update bool, delete bool,
	globalSearch bool,
	globalIterate bool,
	globalRetrieve bool,
	globalCRUD bool) uint64 {
	var p uint64 = 0

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
	if globalCRUD {
		p = setBit(p, 7)
	}

	return p
}

func setBit(n uint64, pos uint16) uint64 {
	n |= 1 << pos
	return n
}

func hasBit(n uint64, pos uint16) bool {
	val := n & (1 << pos)
	return val > 0
}
