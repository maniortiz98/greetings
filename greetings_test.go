package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Juan"
	want := regexp.MustCompile(`\b` + name + `\b`)
	mgs, err := Hello("Juan")
	if !want.MatchString(mgs) || err != nil {
		t.Fatalf(`Hello("Juan") = %q, %v, quiere concidencia para %#q, nil`, mgs, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, quiere un "", error`, msg, err)
	}
}
