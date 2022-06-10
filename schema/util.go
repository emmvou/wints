package schema

import "github.com/emmvou/wints/util"

func IsAdminAtLeast(roles []int) bool {
	return IsRoleAtLeast(roles, AdminLevel)
}

func IsRoleAtLeast(roles []int, level int) bool {
	return util.AnyI(roles, func(r int) bool { return r >= level })
}

func IsRole(roles []int, level int) bool {
	return util.AnyI(roles, func(r int) bool { return r == level })
}
