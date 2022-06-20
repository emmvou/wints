package session

import "github.com/emmvou/wints/schema"

func (s *Session) isAdminAtLeast() bool {
	return schema.IsAdminAtLeast(s.RolesAsLevel())
}
