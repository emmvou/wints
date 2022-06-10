package config

var Cfg Config

//GetParents gets all the parents of a given group
// TODO add error
func GetParents(groups []string, group string) []string {
	if val, ok := Cfg.Internships.Tree[group]; ok {
		groups = append(groups, group)
		if val.Parent != "" {
			return GetParents(groups, val.Parent)
		}
	}
	return groups
}
