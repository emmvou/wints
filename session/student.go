package session

import "github.com/emmvou/wints/schema"

//SetStudentSkippable change the skippable status if the emitter is a major leader at minimum
func (s *Session) SetStudentSkippable(em string, st bool) error {
	if IsAdminAtLeast(s) {
		return s.store.SetStudentSkippable(em, st)
	}
	return ErrPermission
}

//Students lists all the students if the emitter is an admin at least
func (s *Session) Students() (schema.Students, error) { //TODO change when supervisors changed
	students, err := s.store.Students()
	if IsRoleAtLeast(s, schema.HeadLevel) {
		return students, err
	}
	if IsRole(s, schema.SupervisorLevel) {
		return students.Filter(schema.StudentInAllGroups(s.my.AllSubRoles(), s.groups)), err
	}
	return []schema.Student{}, ErrPermission
}

//Student returns the student if the emitter is the student or a watcher
func (s *Session) Student(stu string) (schema.Student, error) {
	if s.Myself(stu) || s.Watching(stu) {
		return s.store.Student(stu)
	}
	return schema.Student{}, ErrPermission
}

//SetAlumni changes the student next position if the emitter is the targetted student,
//the tutor, a member of his jury or a major leader at minimum
func (s *Session) SetAlumni(student string, a schema.Alumni) error {
	if s.Myself(student) || IsRoleAtLeast(s, schema.SupervisorLevel) || s.Tutoring(student) || s.JuryOf(student) {
		return s.store.SetAlumni(student, a)
	}
	return ErrPermission
}

//SetPromotion changes the student group if the emitter is the student themselves or a major leader at least
func (s *Session) SetGroup(student string, g string) error {
	if s.Myself(student) || IsRoleAtLeast(s, schema.SupervisorLevel) {
		return s.store.SetGroup(student, g)
	}
	return ErrPermission
}

//SetMale changes the student gender if the emitter is the student itself, the tutor or an admin at minimum
func (s *Session) SetMale(student string, male bool) error {
	if s.Myself(student) || s.Tutoring(student) || IsRoleAtLeast(s, schema.SupervisorLevel) {
		return s.store.SetMale(student, male)
	}
	return ErrPermission
}
