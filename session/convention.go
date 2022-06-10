package session

import (
	"github.com/emmvou/wints/feeder"
	"github.com/emmvou/wints/schema"
)

//Conventions lists the conventions if the emitter is an admin at minimum
func (s *Session) Conventions() ([]schema.Convention, *feeder.ImportError) {
	if schema.IsAdminAtLeast(s.RolesAsLevel()) {
		return s.conventions.Import()
	}
	ierr := feeder.NewImportError()
	ierr.Fatal = ErrPermission
	return []schema.Convention{}, ierr
}

//Convention returns the convention of a given student if the emitter is the student or at least an admin
func (s *Session) Convention(stu string) (schema.Convention, error) {
	if s.Myself(stu) || schema.IsAdminAtLeast(s.RolesAsLevel()) {
		return s.store.Convention(stu)
	}
	return schema.Convention{}, ErrPermission
}

//SetSupervisor changes the supervisor if the emitter is the student or an admin at minimum
func (s *Session) SetSupervisor(stu string, sup schema.Person) error {
	if s.Myself(stu) || schema.IsAdminAtLeast(s.RolesAsLevel()) {
		return s.store.SetSupervisor(stu, sup)
	}
	return ErrPermission
}

//SetTutor changes the tutor if the emitter is an admin at minimum
func (s *Session) SetTutor(stu string, t string) error {
	if schema.IsAdminAtLeast(s.RolesAsLevel()) {
		return s.store.SetTutor(stu, t)
	}
	return ErrPermission
}

//SetCompany changes the company if the emitter is the student or an admin at minimum
func (s *Session) SetCompany(stu string, c schema.Company) error {
	if s.Myself(stu) || schema.IsAdminAtLeast(s.RolesAsLevel()) {
		return s.store.SetCompany(stu, c)
	}
	return ErrPermission
}

//NewInternship validates the convention if the emitter is an admin at minimum
func (s *Session) NewInternship(c schema.Convention) (schema.Internship, []byte, error) {
	if schema.IsAdminAtLeast(s.RolesAsLevel()) {
		return s.store.NewInternship(c)
	}
	return schema.Internship{}, []byte{}, ErrPermission
}

//Internships list the internships if the emitter is at least a major leader.
//Otherwise, all the internships now tutored by the emitter are removed
// TODO pass group architecture as argument
func (s *Session) Internships() (schema.Internships, error) {
	is, err := s.store.Internships()
	if schema.IsRoleAtLeast(s.RolesAsLevel(), schema.HeadLevel) {
		return is, err
	} else if schema.IsRole(s.RolesAsLevel(), schema.SupervisorLevel) {
		//All the student that are in the major plus the tutored students
		l1 := is.Filter(schema.InGroups(s.my.AllSubRoles()), schema.Tutoring(s.my.Person.Email))
		return l1, err
	} else if schema.IsRole(s.RolesAsLevel(), schema.TutorLevel) {
		return is.Filter(schema.Tutoring(s.my.Person.Email)), err
	}
	//Return the anonymised version
	is.Anonymise()
	return is, nil

}

//Internship returns the internship of the emitter or if the emitter is at least a group leader
func (s *Session) Internship(stu string) (schema.Internship, error) {
	if s.Myself(stu) || s.Watching(stu) {
		return s.store.Internship(stu)
	}
	return schema.Internship{}, ErrPermission
}

//any does any of the given strings validate predicated
func anyS(ss []string, predicate func(string) bool) bool {
	for _, s := range ss {
		if predicate(s) {
			return true
		}
	}
	return false
}
