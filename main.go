package main

import (
	"database/sql"
	"flag"
	"github.com/emmvou/wints/config"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/emmvou/wints/feeder"
	"github.com/emmvou/wints/httpd"
	"github.com/emmvou/wints/jobs"
	"github.com/emmvou/wints/logger"
	"github.com/emmvou/wints/mail"
	"github.com/emmvou/wints/notifier"
	"github.com/emmvou/wints/schema"
	"github.com/emmvou/wints/sqlstore"
	"github.com/robfig/cron"
)

//Version is the running version. Should be set at link time
var Version = "SNAPSHOT"

var not *notifier.Notifier
var store *sqlstore.Store
var mailer mail.Mailer

func confirm(msg string) bool {
	os.Stdout.WriteString(msg + " (y/n) ?")
	b := make([]byte, 1)
	os.Stdin.Read(b)
	ret := string(b) == "y"
	return ret
}

func fatal(msg string, err error) {
	logger.Log("event", "daemon", msg, err)
	st := ""
	if err != nil {
		st = ": " + err.Error()
		log.Fatalln(msg + st)
	} else {
		log.Println(msg + st)
	}
}

func inviteRoot(em string) {
	u := schema.User{
		Person: schema.Person{
			Firstname: "root",
			Lastname:  "root",
			Email:     em,
			Tel:       "n/a",
		},
		Roles: []schema.Role{schema.ROOT},
	}
	token, err := store.NewUser(u.Person, []schema.Role{schema.ROOT})
	fatal("Create root account", err)
	if e := not.InviteRoot(u, u, string(token), err); e != nil {
		//Here, we delete the account as the root account was not aware of the creation
		store.RmUser(u.Person.Email)
		fatal("Invite root", e)
	}
}

func newMailer(fake bool) mail.Mailer {
	if fake {
		return &mail.Fake{WWW: config.Cfg.HTTPd.WWW, Config: config.Cfg.Mailer}
	}
	m, err := mail.NewSMTP(config.Cfg.Mailer, config.Cfg.HTTPd.WWW)
	if err != nil {
		log.Fatalln("SMTP Mailer: " + err.Error())
	}
	return m
}

func newStore() *sqlstore.Store {
	DB, err := sql.Open("postgres", config.Cfg.Db.ConnectionString)
	fatal("Database connexion", err)
	st, _ := sqlstore.NewStore(DB, config.Cfg.Internships)
	return st
}

func newFeeder(not *notifier.Notifier) feeder.Conventions {
	r := feeder.NewHTTPConventionReader(config.Cfg.Feeder.URL, config.Cfg.Feeder.Login, config.Cfg.Feeder.Password)
	r.Encoding = config.Cfg.Feeder.Encoding
	f := feeder.NewCsvConventions(r, config.Cfg.Feeder.Promotions)
	return f
}

func runSpies() {
	c := cron.New()
	if err := c.AddFunc(config.Cfg.Crons.NewsLetters, func() { jobs.MissingReports(store, not) }); err != nil {
		fatal("Unable to cron the missing report scanner", err)
	}
	if err := c.AddFunc(config.Cfg.Crons.Surveys, func() { jobs.MissingSurveys(store, not) }); err != nil {
		fatal("Unable to cron the missing survey scanner", err)
	}
	/*if err := c.AddFunc(Cfg.Crons.Idles, func() { jobs.NeverLogged(store, not) }); err != nil {
		fatal("Unable to cron the idle student scanner", err)
	}*/

	c.Start()
}

func install() {
	if !confirm("This will erase any data. Confirm ") {
		os.Exit(1)
	}
	err := store.Install()
	fatal("Creating tables", err)
}

func insecureServer(listenTo string) {
	s := &http.Server{
		Addr:           listenTo,
		Handler:        http.HandlerFunc(redirectToSecure),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fatal("Insecure listening on "+config.Cfg.HTTPd.InsecureListen, nil)
	err := s.ListenAndServe()
	fatal("Stop insecure listening on "+config.Cfg.HTTPd.InsecureListen, err)
}

func redirectToSecure(w http.ResponseWriter, req *http.Request) {
	logger.Log("event", "insecure", "redirection to "+config.Cfg.HTTPd.WWW+req.RequestURI, nil)
	http.Redirect(w, req, config.Cfg.HTTPd.WWW+req.RequestURI, http.StatusMovedPermanently)
}

func main() {

	makeRoot := flag.String("new-root", "", "Invite a root user")
	fakeMailer := flag.Bool("fake-mailer", false, "Don't send emails. Print them out stdout")
	installStore := flag.Bool("install-db", false, "install the database")
	conf := flag.String("conf", "wints.conf", "Wints configuration file")
	flag.Parse()

	_, err := toml.DecodeFile(*conf, &config.Cfg)
	if err != nil {
		log.Fatalf("reading configuration '%s': %s\n", *conf, err.Error())
	}

	err = logger.SetRoot(config.Cfg.Journal.Path)
	if len(config.Cfg.Journal.Key) > 0 {
		logger.Trace(config.Cfg.Journal.Key)
	}
	fatal("Initiating the logger", err)
	fatal("Running Version '"+Version+"'", nil)
	config.Cfg.Internships.Version = Version
	mailer = newMailer(*fakeMailer)
	not = notifier.New(mailer)

	store = newStore()
	if *installStore {
		install()
		os.Exit(0)
	}

	_, err = store.Internships()
	fatal("Database communication", err)

	if len(*makeRoot) > 0 {
		inviteRoot(*makeRoot)
		os.Exit(0)
	}

	runSpies()
	conventions := newFeeder(not)

	//Insecure listen that only redirect
	go insecureServer(config.Cfg.HTTPd.InsecureListen)
	fatal("Listening on "+config.Cfg.HTTPd.WWW, nil)
	httpd_ := httpd.NewHTTPd(not, store, conventions)
	err = httpd_.Listen()
	fatal("Stop listening on "+config.Cfg.HTTPd.WWW, err)
}
