package sqlstore

import (
	"database/sql"
	"errors"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/emmvou/wints/schema"
	"github.com/lib/pq"
)

var (
	selectUser                   = "select firstname, lastname, email, tel, lastVisit from users where email=$1"
	insertUser                   = "insert into users(firstname, lastname, tel, email, role, password) values ($1,$2,$3,$4,$5,$6)" // TODO
	startPasswordRenewal         = "insert into password_renewal(email,token) values($1,$2)"
	updateLastVisit              = "update users set lastVisit=$1 where email=$2"
	updateUserProfile            = "update users set firstname=$1, lastname=$2, tel=$3 where email=$4"
	updateUserRole               = "update users set role=$2 where email=$1 and role != 'student'" //cannot change the role of a student // TODO
	updateUserPassword           = "update users set password=$2 where email=$1"
	updateEmail                  = "update users set email=$2 where email=$1"
	deletePasswordRenewalRequest = "delete from password_renewal where email=$1"
	deleteUser                   = "DELETE FROM users where email=$1"
	allUsers                     = "select u.firstname, u.lastname, u.email, u.tel, u.lastVisit from users u"
	allUserRoles                 = "select u.email, r.role from users u join userroles ur on u.email=ur.user join roles r on ur.role=r.id"
	allRoles                     = "select role from roles"
	selectPassword               = "select password from users where email=$1"
	emailFromRenewableToken      = "select email from password_renewal where token=$1"
	replaceTutorInConventions    = "update conventions set tutor=$2 where tutor=$1"
	replaceJuryInDefenses        = "update defenseJuries set jury=$2 where jury=$1"
	selectAlias                  = "select real from aliases where email=$1"
	insertAlias                  = "insert into aliases(email,real) values($1,$2)"
)

//addUser add the given user.
//Every strings are turned into their lower case version
func (s *Store) addUser(tx *TxErr, u schema.User) {
	if !strings.Contains(u.Person.Email, "@") {
		tx.err = schema.ErrInvalidEmail
		return
	}
	//Check if aliased name (this email corresponds to another real email)
	var em string
	err := tx.QueryRow(selectAlias, u.Person.Email).Scan(&em)
	if err == nil {
		//There is already a user. The proposed email was just an alias
		tx.err = schema.ErrUserExists
	} else {
		if err != sql.ErrNoRows {
			tx.err = err
		}
	}
	tx.Exec(insertUser,
		u.Person.Firstname,
		u.Person.Lastname,
		u.Person.Tel,
		u.Person.Email,
		u.Role.String(), // TODO
		randomBytes(32),
	)
	tx.Exec(insertAlias, u.Person.Email, u.Person.Email)
}

//Visit writes the current time for the given user
func (s *Store) Visit(u string) error {
	return s.singleUpdate(updateLastVisit, schema.ErrUnknownUser, time.Now(), u)
}

func scanUser(row *sql.Rows) (schema.User, error) {
	u := schema.User{
		Person: schema.Person{},
		Roles:  []schema.Role{},
	}
	var last pq.NullTime
	var roles []string
	err := row.Scan(
		&u.Person.Firstname,
		&u.Person.Lastname,
		&u.Person.Email,
		&u.Person.Tel,
		&last,
		&roles,
	)
	u.LastVisit = nullableTime(last)
	return u, err
}

func scanUserRole(row *sql.Rows) (schema.UserRole, error) {
	ur := schema.UserRole{
		Role: schema.Role(""),
	}
	err := row.Scan(
		&ur.User,
		&ur.Role,
	)
	return ur, err
}

func scanRole(rows *sql.Rows) (schema.Role, error) {
	ur := schema.Role("")
	var role string
	err := rows.Scan(
		&role,
	)
	ur = schema.Role(role)
	return ur, err
}

func (s *Store) user(email string) (schema.User, error) {
	st := s.stmt(selectUser)
	rows, err := st.Query(email)
	if err != nil {
		return schema.User{}, err
	}
	defer rows.Close()
	if !rows.Next() {
		return schema.User{}, schema.ErrUnknownUser
	}
	return scanUser(rows)
}

//User returns the given user account
func (s *Store) User(email string) (schema.User, error) {
	user, err := s.user(email)
	if err != nil {
		return schema.User{}, err
	}
	roles, err := s.Roles()
	if err != nil {
		return schema.User{}, err
	}
	userRoles, err := s.userRoles()
	if err != nil {
		return schema.User{}, err
	}
	// yes, this is a bit ugly, but it's the only way to do it
	for _, ur := range userRoles {
		if ur.User == user.Person.Email {
			for _, r := range roles {
				if r.String() == ur.Role.String() {
					user.Roles = append(user.Roles, r)
				}
			}
		}
	}

	return user, nil
}

func (s *Store) users() ([]schema.User, error) {
	var users []schema.User
	st := s.stmt(allUsers)
	rows, err := st.Query()
	if err != nil {
		return users, err
	}
	defer rows.Close()
	for rows.Next() {
		u, err := scanUser(rows)
		if err != nil {
			return users, err
		}
		users = append(users, u)
	}
	return users, nil
}

//Users list all the registered users
func (s *Store) Users() ([]schema.User, error) {
	users, err := s.users()
	if err != nil {
		return []schema.User{}, err
	}
	roles, err := s.Roles()
	if err != nil {
		return []schema.User{}, err
	}
	userRoles, err := s.userRoles()
	if err != nil {
		return []schema.User{}, err
	}
	// yes, this is a bit ugly, but it's the only way to do it
	for _, u := range users {
		for _, ur := range userRoles {
			if ur.User == u.Person.Email {
				for _, r := range roles {
					if r.String() == ur.Role.String() {
						u.Roles = append(u.Roles, r)
					}
				}
			}
		}
	}

	return users, nil
}

//Users list all the registered users
func (s *Store) userRoles() ([]schema.UserRole, error) {
	var users []schema.UserRole
	st := s.stmt(allUserRoles)
	rows, err := st.Query()
	if err != nil {
		return users, err
	}
	defer rows.Close()
	for rows.Next() {
		u, err := scanUserRole(rows)
		if err != nil {
			return users, err
		}
		users = append(users, u)
	}
	return users, nil
}

//Roles list all possible roles
func (s *Store) Roles() ([]schema.Role, error) {
	var roles []schema.Role
	st := s.stmt(allRoles)
	rows, err := st.Query()
	if err != nil {
		return roles, err
	}
	defer rows.Close()
	for rows.Next() {
		r, err := scanRole(rows)
		if err != nil {
			return roles, err
		}
		roles = append(roles, r)
	}

	return roles, nil
}

//SetUserPerson changes the user profile if exists
func (s *Store) SetUserPerson(p schema.Person) error {
	return s.singleUpdate(updateUserProfile, schema.ErrUnknownUser, p.Firstname, p.Lastname, p.Tel, p.Email)
}

//SetUserRole updates the user privilege
func (s *Store) SetUserRole(email string, priv schema.Role) error {
	return s.singleUpdate(updateUserRole, schema.ErrUnknownUser, email, priv) // TODO
}

//ResetPassword starts a reset procedure
func (s *Store) ResetPassword(email string) ([]byte, error) {
	token := randomBytes(32)

	//If exists & not outdate, resend the token
	//otherwise new token
	tx := newTxErr(s.db)
	tx.Exec(deletePasswordRenewalRequest, email)
	//In case a request already exists
	tx.Exec(startPasswordRenewal, email, token)
	return token, tx.Done()
}

//NewPassword commits a password renewall request.
//From a request token and a new password, it returns upon success the target user email
func (s *Store) NewPassword(token, newP []byte) (string, error) {
	if len(newP) < 8 {
		return "", schema.ErrPasswordTooShort
	}
	hash, err := bcrypt.GenerateFromPassword(newP, bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	var email string
	tx := newTxErr(s.db)
	tx.err = tx.QueryRow(emailFromRenewableToken, token).Scan(&email)
	tx.err = noRowsTo(tx.err, schema.ErrNoPendingRequests)
	tx.Update(updateUserPassword, email, hash)
	//no need to check updated rows as it is sure the user exists in the tx context
	nb := tx.Update(deletePasswordRenewalRequest, email)
	if tx.err == nil && nb != 1 {
		tx.err = errors.New("Unable to clean the password renewable request of " + email)
	}
	return email, tx.Done()
}

//NewUser add a user
//Basically, calls addUser
func (s *Store) NewUser(p schema.Person, roles []schema.Role) ([]byte, error) {
	if !validEmail(p.Email) {
		return []byte{}, schema.ErrInvalidEmail
	}
	token := randomBytes(32)
	tx := newTxErr(s.db)
	nb := tx.Update(insertUser, p.Firstname, p.Lastname, p.Tel, p.Email, roles.String(), randomBytes(32)) // TODO
	if nb == 0 {
		tx.err = schema.ErrUserExists
	}
	tx.Exec(startPasswordRenewal, p.Email, token)
	return token, tx.Done()
}

//RmUser removes a user from the database
func (s *Store) RmUser(email string) error {
	_, err := s.Internship(email)
	if err == nil {
		return schema.ErrInternshipExists
	}
	return s.singleUpdate(deleteUser, schema.ErrUnknownUser, email)
}

//SetEmail change a user email to another
func (s *Store) SetEmail(old, now string) error {
	if !validEmail(now) {
		return schema.ErrInvalidEmail
	}
	//Create an alias to remember the email
	tx := newTxErr(s.db)
	tx.Update(updateEmail, old, now)
	//Alias after the user because it does not exists otherwise
	tx.Exec(insertAlias, old, now)
	return tx.Done()
}

//ReplaceUserWith the account referred by src by the account referred by dst
func (s *Store) ReplaceUserWith(src, dst string) error {
	tx := newTxErr(s.db)
	tx.Update(replaceTutorInConventions, src, dst)
	tx.Update(replaceJuryInDefenses, src, dst)
	tx.Update(deleteUser, src)
	return tx.Done()
}

func validEmail(em string) bool {
	if !strings.Contains(em, "@") {
		return false
	}
	for _, c := range []string{";", ",", " "} {
		if strings.Contains(em, c) {
			return false
		}
	}
	return true
}
