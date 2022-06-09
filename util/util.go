package util

import (
	"github.com/emmvou/wints/config"
	"github.com/emmvou/wints/schema"
)

var Cfg config.Config

func IsAdminAtLeast(roles []int) bool {
	return IsRoleAtLeast(roles, schema.AdminLevel)
}

func IsRoleAtLeast(roles []int, level int) bool {
	return anyI(roles, func(r int) bool { return r >= level })
}

//any item from roles validates predicate
func anyI(roles []int, predicate func(r int) bool) bool {
	for _, s := range roles {
		if predicate(s) {
			return true
		}
	}
	return false
}

func IsRole(roles []int, level int) bool {
	return anyI(roles, func(r int) bool { return r == level })
}

// TODO add error
func GetParents(groups []string, group string) []string {
	if val, ok := Cfg.Internships.Groups[group]; ok {
		groups = append(groups, group)
		if val.Parent != "" {
			return GetParents(groups, val.Parent)
		}
	}
	return groups
}

func RemoveDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
