package sqlstore

import "github.com/emmvou/wints/schema"

func rolesToInts(roles []schema.Role) []int {
	var res []int
	for _, r := range roles {
		res = append(res, r.Level())
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
