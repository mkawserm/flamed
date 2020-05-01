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

	return (p & 1 << (0)) == 1
}

func HasWritePermission(ac *pb.FlameAccessControl) bool {
	if ac == nil {
		return false
	}

	if len(ac.Permission) == 0 {
		return false
	}

	var p = ac.Permission[0]

	return (p & 1 << (1)) == 1
}

func HasUpdatePermission(ac *pb.FlameAccessControl) bool {
	if ac == nil {
		return false
	}

	if len(ac.Permission) == 0 {
		return false
	}

	var p = ac.Permission[0]

	return (p & 1 << (2)) == 1
}

func HasDeletePermission(ac *pb.FlameAccessControl) bool {
	if ac == nil {
		return false
	}

	if len(ac.Permission) == 0 {
		return false
	}

	var p = ac.Permission[0]

	return (p & 1 << (3)) == 1
}
