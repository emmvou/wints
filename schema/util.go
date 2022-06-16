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

func rolesAsLevel(user *User) []int {
	var roles []int
	for _, r := range user.Roles {
		roles = append(roles, r.Level())
	}
	return roles
}

func IsStudent(user *User) bool {
	return IsRole(rolesAsLevel(user), StudentLevel)
}
