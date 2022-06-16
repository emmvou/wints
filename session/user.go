package session

import "github.com/emmvou/wints/schema"

//RmUser removes an account if the emitter is at least an admin and not himself
func (s *Session) RmUser(email string) error {
	if schema.IsAdminAtLeast(s.RolesAsLevel()) && !s.Myself(email) {
		return s.store.RmUser(email)
	}
	return ErrPermission
}

//Users lists all the users if the emitter is at least an admin
func (s *Session) Users() ([]schema.User, error) {
	if schema.IsAdminAtLeast(s.RolesAsLevel()) {
		return s.store.Users()
	}
	return []schema.User{}, ErrPermission
}

//User returns a given user if the emitter is himself or an admin
func (s *Session) User(em string) (schema.User, error) {
	if s.Myself(em) || schema.IsAdminAtLeast(s.RolesAsLevel()) {
		return s.store.User(em)
	}
	return schema.User{}, ErrPermission
}

//SetUserPerson set the user profile if the emitter is the targeted user
func (s *Session) SetUserPerson(p schema.Person) error {
	if s.Myself(p.Email) || schema.IsAdminAtLeast(s.RolesAsLevel()) {
		return s.store.SetUserPerson(p)
	}
	return ErrPermission
}

//SetUserRole changes the user privileges if the emitter is an admin at minimum and not himself
func (s *Session) SetUserRole(email string, priv []schema.Role) error {
	if schema.IsAdminAtLeast(s.RolesAsLevel()) && !s.Myself(email) {
		return s.store.SetUserRole(email, priv)
	}
	return ErrPermission
}

//NewUser creates a new user account if the emitter is an admin at least
func (s *Session) NewUser(p schema.Person, roles []schema.Role) ([]byte, error) {
	if schema.IsAdminAtLeast(s.RolesAsLevel()) {
		return s.store.NewUser(p, roles)
	}
	return []byte{}, ErrPermission
}

//NewStudent creates a new student account if the emitter is an admin at least
func (s *Session) NewStudent(p schema.Person, group string, male bool) error {
	if schema.IsAdminAtLeast(s.RolesAsLevel()) {
		return s.store.NewStudent(p, group, male)
	}
	return ErrPermission
}

//ReplaceUserWith allowed if the emitter is an admin at least
func (s *Session) ReplaceUserWith(src, dst string) error {
	if schema.IsAdminAtLeast(s.RolesAsLevel()) {
		return s.store.ReplaceUserWith(src, dst)
	}
	return ErrPermission
}

//SetEmail change the person email if the emitter is at least an admin
func (s *Session) SetEmail(old, cur string) error {
	if schema.IsAdminAtLeast(s.RolesAsLevel()) {
		return s.store.SetEmail(old, cur)
	}
	return ErrPermission
}
