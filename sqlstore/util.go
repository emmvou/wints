package sqlstore

import "github.com/emmvou/wints/schema"

func rolesToInts(roles []schema.Role) []int {
	var res []int
	for _, r := range roles {
		res = append(res, r.Level())
	}
	return res
}

func rolesToStrings(roles []schema.Role) []string {
	var res []string
	for _, r := range roles {
		res = append(res, r.String())
	}
	return res
}

func removeDuplicateRoles(strSlice []schema.Role) []schema.Role {
	allKeys := make(map[schema.Role]bool)
	list := []schema.Role{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func isOverlapRoles(a, b []schema.Role) bool {
	m := make(map[schema.Role]bool)
	for _, r := range a {
		m[r] = true
	}
	for _, r := range b {
		if _, ok := m[r]; ok {
			return true
		}
	}
	return false
}

//are all the elements of a in b?
//ignores duplicates
func isSliceInSlice(a, b []schema.Role) bool {
	m := make(map[schema.Role]bool)
	for _, r := range b {
		m[r] = true
	}
	for _, r := range a {
		if _, ok := m[r]; !ok {
			return false
		}
	}
	return true
}
