//Package config aggregates the necessary material to configure the wints daemon
package config

import "github.com/emmvou/wints/mail"

var (
	//DateTimeLayout expresses the expected format for a date + time
	DateTimeLayout = "02/01/2006 15:04"
	//DateLayout expresses the expected format for a date
	DateLayout = "02/01/2006"
)

//Feeder configures the feeder than scan conventions
type Feeder struct {
	Login      string
	Password   string
	URL        string
	Frequency  Duration
	Promotions []string //TODO replace with groups
	Encoding   string
}

//Db configures the database connection string
type Db struct {
	ConnectionString string
}

//Rest configures the rest service
type Rest struct {
	SessionLifeTime        Duration
	RenewalRequestLifetime Duration
	//The endpoints prefix
	Prefix string
}

//HTTPd configures the http daemon
type HTTPd struct {
	InsecureListen string
	WWW            string
	Assets         string
	Listen         string
	Certificate    string
	PrivateKey     string
	Rest           Rest
}

//Journal configures the logging system
type Journal struct {
	Path string
	Key  string
}

//Internships declare the internship organization
type Internships struct {
	Groups      map[string]Group
	Reports     []Report
	Surveys     []Survey
	LatePenalty int
	Version     string
}

type Group struct {
	Name   string
	Parent string
}

func contains(s []string, v string) bool {
	for _, x := range s {
		if x == v {
			return true
		}
	}
	return false
}

func (i Internships) ValidGroup(g string) bool {
	_, ok := i.Groups[g]
	return ok
}

//Report configures a report definition
type Report struct {
	Kind     string
	Delivery Deadline
	Review   Duration
	Reminder Duration
	Grade    bool
}

//Survey configures a survey definition
type Survey struct {
	Invitation Deadline
	Deadline   Duration
	Kind       string
}

//Crons list the waiting periods for some period tasks
type Crons struct {
	//NewsLetters is the waiting time between two news letters
	NewsLetters string
	//Surveys is the waiting time between two scan for missing surveys
	Surveys string
	//Idles is the waiting time between two scan for missing student connections
	Idles string
}

//Config aggregates all the subcomponents configuration parameters
type Config struct {
	Feeder      Feeder
	Db          Db
	Mailer      mail.Config
	HTTPd       HTTPd
	Groups      map[string]Group
	Journal     Journal
	Internships Internships
	Crons       Crons
}
