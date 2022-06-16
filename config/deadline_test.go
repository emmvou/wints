package config

import (
	"testing"
	"time"

	"github.com/BurntSushi/toml"
)

var (
	relative = `deadline = "1440h"`
	absolute = `deadline = "01/09/2015 15:00"`
)

type D struct {
	Deadline Deadline
}

func TestRelativeDeadline(t *testing.T) {
	var d D
	if _, err := toml.Decode(relative, &d); err != nil {
		t.Error(err.Error())
	}
	expected_, _ := time.ParseDuration("1440h")
	if *d.Deadline.relative != expected_ {
		t.Errorf("Got %s but expected %s\n", *d.Deadline.relative, expected_)
	}
	if d.Deadline.absolute != nil {
		t.Error("absolute value should be nil")
	}
	now := time.Now()
	v := d.Deadline.Value(now)
	if v != now.Add(expected_) {
		t.Errorf("Got %s but expected %s\n", v, now.Add(expected_))
	}
}

func TestAbsoluteDeadline(t *testing.T) {
	var d D
	if _, err := toml.Decode(absolute, &d); err != nil {
		t.Error(err.Error())
	}
	expected_, err := time.Parse(DateTimeLayout, "01/09/2015 15:00")
	if err != nil {
		t.Fatalf("%s", err.Error())
	}
	if *d.Deadline.absolute != expected_ {
		t.Errorf("Got %s but expected %s\n", *d.Deadline.absolute, expected_)
	}
	if d.Deadline.relative != nil {
		t.Error("absolute value should be nil")
	}
	v := d.Deadline.Value(time.Now())
	if v != expected_ {
		t.Errorf("Got %s but expected %s\n", v, expected_)
	}
}
