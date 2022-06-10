//Package session provides the material to implement the security layer.
//A session restrict the operation that are authorized by the session owner, depending on its identity or his role
package session

import (
	"errors"
	"github.com/emmvou/wints/util"

	"github.com/emmvou/wints/config"
	"github.com/emmvou/wints/feeder"
	"github.com/emmvou/wints/schema"
	"github.com/emmvou/wints/sqlstore"
)

var (
	//ErrPermission indicates an operation that is not permitted
	ErrPermission = errors.New("Permission denied")
)

//Session restricts the operation that can be executed by the current user with regards
//to its role and or relationships
type Session struct {
	my          schema.User
	store       *sqlstore.Store
	conventions feeder.Conventions
	groups      map[string]*config.Group
}

//AnonSession creates a session that is not attached to a particular user
func AnonSession(store *sqlstore.Store) Session {
	return Session{store: store}
}

//NewSession creates a new session
func NewSession(u schema.User, store *sqlstore.Store, conventions feeder.Conventions, groups map[string]*config.Group) Session {
	return Session{my: u, store: store, conventions: conventions, groups: groups}
}

//RmSession delete the session if the emitter is the session owner or at least an admin
func (s *Session) RmSession(em string) error {
	if s.Myself(em) || schema.IsAdminAtLeast(s.RolesAsLevel()) {
		return s.store.RmSession(em)
	}
	return ErrPermission
}

//Me returns the session emitter
func (s *Session) Me() schema.User {
	return s.my
}

func (s *Session) Roles() []schema.Role {
	return s.my.Roles
}

func (s *Session) RolesAsLevel() []int {
	var roles []int
	for _, r := range s.my.Roles {
		roles = append(roles, r.Level())
	}
	return roles
}

//Myself checks if the given email matches the session one
func (s *Session) Myself(email string) bool {
	return s.my.Person.Email == email
}

//Tutoring checks if the session owner is a tutor of the given student
func (s *Session) Tutoring(student string) bool {
	c, err := s.store.Convention(student)
	if err != nil {
		return false
	}
	return c.Tutor.Person.Email == s.my.Person.Email
}

//Watching checks if the student is in the major I am the leader of, a head or not or
//the student tutor
// TODO inject group arborescence
func (s *Session) Watching(student string) bool {
	//watching
	if schema.IsRoleAtLeast(s.RolesAsLevel(), schema.HeadLevel) {
		return true
	}
	c, err := s.store.Convention(student)
	if err != nil {
		return false
	}
	//watching
	if s.InMyGroups(student) {
		return true
	}
	//tutoring
	return s.my.Person.Email == c.Tutor.Person.Email
}

//JuryOf checks if I am in a jury for a defense.
//That if indeed I am in the jury, or an admin
func (s *Session) JuryOf(student string) bool {
	if schema.IsAdminAtLeast(s.RolesAsLevel()) {
		return true
	}
	def, err := s.store.Defense(student)
	if err != nil {
		return false
	}
	mySession, err := s.store.DefenseSession(def.Room, def.SessionId)
	if err != nil {
		return false
	}
	return mySession.InJury(s.my.Person.Email)
}

//InMyGroups checks if the student is in the group or a subgroup I am the leader of
func (s *Session) InMyGroups(student string) bool {
	//call InGroup for each group I am the leader of + remove redundancies
	stu, err := s.store.Student(student)
	if err != nil {
		return false
	}
	//get all parents of the group
	var groups []string
	for _, group := range s.my.AllSubRoles() {
		groups = config.GetParents(groups, group)
	}
	//remove redundancies
	groups = util.RemoveDuplicateStr(groups)

	return util.StringInSlice(stu.Group, groups)
}
