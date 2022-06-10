package main

import (
	"database/sql"
	"flag"
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
	"github.com/emmvou/wints/util"
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
		return &mail.Fake{WWW: util.Cfg.HTTPd.WWW, Config: util.Cfg.Mailer}
	}
	m, err := mail.NewSMTP(util.Cfg.Mailer, util.Cfg.HTTPd.WWW)
	if err != nil {
		log.Fatalln("SMTP Mailer: " + err.Error())
	}
	return m
}

func newStore() *sqlstore.Store {
	DB, err := sql.Open("postgres", util.Cfg.Db.ConnectionString)
	fatal("Database connexion", err)
	st, _ := sqlstore.NewStore(DB, util.Cfg.Internships)
	return st
}

func newFeeder(not *notifier.Notifier) feeder.Conventions {
	r := feeder.NewHTTPConventionReader(util.Cfg.Feeder.URL, util.Cfg.Feeder.Login, util.Cfg.Feeder.Password)
	r.Encoding = util.Cfg.Feeder.Encoding
	f := feeder.NewCsvConventions(r, util.Cfg.Feeder.Promotions)
	return f
}

func runSpies() {
	c := cron.New()
	if err := c.AddFunc(util.Cfg.Crons.NewsLetters, func() { jobs.MissingReports(store, not) }); err != nil {
		fatal("Unable to cron the missing report scanner", err)
	}
	if err := c.AddFunc(util.Cfg.Crons.Surveys, func() { jobs.MissingSurveys(store, not) }); err != nil {
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
	fatal("Insecure listening on "+util.Cfg.HTTPd.InsecureListen, nil)
	err := s.ListenAndServe()
	fatal("Stop insecure listening on "+util.Cfg.HTTPd.InsecureListen, err)
}

func redirectToSecure(w http.ResponseWriter, req *http.Request) {
	logger.Log("event", "insecure", "redirection to "+util.Cfg.HTTPd.WWW+req.RequestURI, nil)
	http.Redirect(w, req, util.Cfg.HTTPd.WWW+req.RequestURI, http.StatusMovedPermanently)
}

func main() {

	makeRoot := flag.String("new-root", "", "Invite a root user")
	fakeMailer := flag.Bool("fake-mailer", false, "Don't send emails. Print them out stdout")
	installStore := flag.Bool("install-db", false, "install the database")
	conf := flag.String("conf", "wints.conf", "Wints configuration file")
	flag.Parse()

	_, err := toml.DecodeFile(*conf, &util.Cfg)
	if err != nil {
		log.Fatalf("reading configuration '%s': %s\n", *conf, err.Error())
	}

	err = logger.SetRoot(util.Cfg.Journal.Path)
	if len(util.Cfg.Journal.Key) > 0 {
		logger.Trace(util.Cfg.Journal.Key)
	}
	fatal("Initiating the logger", err)
	fatal("Running Version '"+Version+"'", nil)
	util.Cfg.Internships.Version = Version
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
	go insecureServer(util.Cfg.HTTPd.InsecureListen)
	fatal("Listening on "+util.Cfg.HTTPd.WWW, nil)
	httpd := httpd.NewHTTPd(not, store, conventions)
	err = httpd.Listen()
	fatal("Stop listening on "+util.Cfg.HTTPd.WWW, err)
}
