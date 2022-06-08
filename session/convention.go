package session

import (
	"github.com/emmvou/wints/feeder"
	"github.com/emmvou/wints/schema"
)

//Conventions lists the conventions if the emitter is an admin at minimum
func (s *Session) Conventions() ([]schema.Convention, *feeder.ImportError) {
	if IsAdminAtLeast(s) {
		return s.conventions.Import()
	}
	ierr := feeder.NewImportError()
	ierr.Fatal = ErrPermission
	return []schema.Convention{}, ierr
}

//Convention returns the convention of a given student if the emitter is the student or at least an admin
func (s *Session) Convention(stu string) (schema.Convention, error) {
	if s.Myself(stu) || IsAdminAtLeast(s) {
		return s.store.Convention(stu)
	}
	return schema.Convention{}, ErrPermission
}

//SetSupervisor changes the supervisor if the emitter is the student or an admin at minimum
func (s *Session) SetSupervisor(stu string, sup schema.Person) error {
	if s.Myself(stu) || IsAdminAtLeast(s) {
		return s.store.SetSupervisor(stu, sup)
	}
	return ErrPermission
}

//SetTutor changes the tutor if the emitter is an admin at minimum
func (s *Session) SetTutor(stu string, t string) error {
	if IsAdminAtLeast(s) {
		return s.store.SetTutor(stu, t)
	}
	return ErrPermission
}

//SetCompany changes the company if the emitter is the student or an admin at minimum
func (s *Session) SetCompany(stu string, c schema.Company) error {
	if s.Myself(stu) || IsAdminAtLeast(s) {
		return s.store.SetCompany(stu, c)
	}
	return ErrPermission
}

//NewInternship validates the convention if the emitter is an admin at minimum
func (s *Session) NewInternship(c schema.Convention) (schema.Internship, []byte, error) {
	if IsAdminAtLeast(s) {
		return s.store.NewInternship(c)
	}
	return schema.Internship{}, []byte{}, ErrPermission
}

// TODO pass group architecture as argument
//Internships list the internships if the emitter is at least a major leader.
//Otherwise, all the internships now tutored by the emitter are removed
func (s *Session) Internships() (schema.Internships, error) {
	is, err := s.store.Internships()
	if IsRoleAtLeast(s, schema.HeadLevel) {
		return is, err
	} else if IsRole(s, schema.SupervisorLevel) {
		//All the student that are in the major plus the tutored students
		l1 := is.Filter(schema.InGroups(s.my.AllSubRoles(), s.groups), schema.Tutoring(s.my.Person.Email))
		return l1, err
	} else if IsRole(s, schema.TutorLevel) {
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

//any does any of the given strings validate predicated
func anyR(ss []schema.Role, predicate func(int) bool) bool {
	for _, s := range ss {
		if predicate(s.Level()) {
			return true
		}
	}
	return false
}

func IsAdminAtLeast(s *Session) bool {
	return IsRoleAtLeast(s, schema.AdminLevel)
}

func IsRoleAtLeast(s *Session, r int) bool {
	return anyR(s.Roles(), func(rr int) bool { return rr >= r })
}

func IsRole(s *Session, r int) bool {
	return anyR(s.Roles(), func(rr int) bool { return rr == r })
}
