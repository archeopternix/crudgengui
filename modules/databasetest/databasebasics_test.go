{{define "databasebasicstest" -}}
// Package mockdatabase contains structures and function for mock database access
// Generated code - do not modify it will be overwritten!!
// Time: {{.TimeStamp}}
package database

import (
	"math/rand"
	"strings"
	"testing"
)

{{range .Entities}}
var {{.Name | lowercase}}db *{{.Name}}Repo
{{end}}

const text = "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam" +
	" nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua." +
	" At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren," +
	" no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet," +
	" consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat," +
	" sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum." +
	" Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet."

var email = []string{"email@example.com",
	"firstname.lastname@example.com",
	"email@subdomain.example.com",
	"firstname+lastname@example.com",
	"1234567890@example.com",
	"email@example-one.com",
	"email@example.name",
	"email@example.museum",
	"email@example.co.jp",
	"firstname-lastname@example.com",
}

func getText(maxlength int) string {
	rand.Seed(42)
	if maxlength > (len(text) + 3) {
		maxlength = len(text) - 5
	}
	le := rand.Intn((len(text) - maxlength))
	str := text[le:(le + maxlength)]
	return strings.TrimSpace(str)
}

func getEmail() string {
	le := rand.Intn(len(email))
	return email[le]
}

func TestGetText(t *testing.T) {
	t.Logf("get long text: '%s'", getText(112))
	t.Logf("get email address: '%s'", getEmail())
}

{{end}}