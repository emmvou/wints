package schema

import "errors"

var (
	//ErrRoleStudent declares the role cannot be changed when it is student
	ErrRoleStudent = errors.New("cannot change role when role is student")
	//ErrUnknownStudent declares the student is unknown
	ErrUnknownStudent = errors.New("unknown student")
	//ErrRolesConflict declares the user cannot have multiple roles when one of them is student
	ErrRolesConflict = errors.New("user cannot have multiple roles including student")
	//ErrOverlapRoles declare user can have a role only once
	ErrOverlapRoles = errors.New("user cannot have the same role more than once")
	//ErrRoleNotPresent declares the user does not have the desired role
	ErrRoleNotPresent = errors.New("role not present")
	//ErrUnknownConvention declares the convention is unknown
	ErrUnknownConvention = errors.New("no convention associated to this student")
	//ErrStudentExists declares the student already exists
	ErrStudentExists = errors.New("student already exists")
	//ErrReportExists declares the report already exists
	ErrReportExists = errors.New("report already exists")
	//ErrUnknownReport declares the report does not exist
	ErrUnknownReport = errors.New("unknown report or student")
	//ErrInvalidGrade declares the grade is not between in 0 and 20
	ErrInvalidGrade = errors.New("the grade must be between 0 and 20 (inclusive)")
	//ErrReportConflict declares a report has not been uploaded
	ErrReportConflict = errors.New("the report has not been uploaded")
	//ErrInternshipExists declares the internship already exists
	ErrInternshipExists = errors.New("internship already exists")
	//ErrUnknownInternship declares the internship does not exists
	ErrUnknownInternship = errors.New("unknown internship")
	//ErrUserExists declares the user already exists
	ErrUserExists = errors.New("user already exists")
	//ErrUnknownUser declares the user is unknown
	ErrUnknownUser = errors.New("the email does not match a registered user")
	//ErrUserTutoring declares the user cannot be removed as it is tutoring students
	ErrUserTutoring = errors.New("the user is tutoring students")
	//ErrCredentials declares invalid credentials
	ErrCredentials = errors.New("incorrect password")
	//ErrPasswordTooShort declares the stated password is too short
	ErrPasswordTooShort = errors.New("password too short (8 chars. min)")
	//ErrNoPendingRequests declares there is no password renewal request
	ErrNoPendingRequests = errors.New("no pending reset request. You might use a bad or an expired reset token")
	//ErrInvalidPeriod declared the internship period is incorrect
	ErrInvalidPeriod = errors.New("invalid internship period")
	//ErrConventionExists declares a convention for the student already exists
	ErrConventionExists = errors.New("convention already scanned")
	//ErrInvalidGroup declares the group is not supported
	ErrInvalidGroup = errors.New("unknown group")
	//ErrDeadlinePassed declares the deadline for a report passed
	ErrDeadlinePassed = errors.New("deadline passed")
	//ErrGradedReport declares the report is already graded
	ErrGradedReport = errors.New("report already graded")
	//ErrSessionExpired declares an expired session
	ErrSessionExpired = errors.New("session expired")
	//ErrInvalidToken declares an invalid session token
	ErrInvalidToken = errors.New("invalid session")
	//ErrUnknownSurvey declares the survey does not exist
	ErrUnknownSurvey = errors.New("unknown survey or student")
	//ErrSurveyUploaded declares the survey has already been uploaded
	ErrSurveyUploaded = errors.New("survey already fullfilled")
	//ErrInvalidSurvey declares the answers are invalid
	ErrInvalidSurvey = errors.New("invalid answers")
	//ErrUnknownAlumni declares there is no alumni information for the students
	ErrUnknownAlumni = errors.New("no informations for future alumni")
	//ErrInvalidAlumniEmail declares the email cannot be used for alumni
	ErrInvalidAlumniEmail = errors.New("invalid email. It must not be served by polytech' or unice")
	//ErrInvalidEmail declares the email is invalid
	ErrInvalidEmail = errors.New("invalid email")
	//ErrUnknownDefense declares the defense is unknown
	ErrUnknownDefense = errors.New("unknown defense")

	//ErrDefenseSessionConflit declares a session that is in conflict with another
	ErrDefenseSessionConflit = errors.New("there is already a session for that slot")
	ErrDefenseExists         = errors.New("the defense is already planned")
	ErrDefenseConflict       = errors.New("a defense is already planned for that slot")
	ErrDefenseJuryConflict   = errors.New("the teacher is already in a jury for that period")
)
