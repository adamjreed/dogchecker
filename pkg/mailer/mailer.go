package mailer

import (
	"bytes"
	_ "embed"
	"github.com/friendsofgo/errors"
	"html/template"
)

type Config struct {
	ApiKey      string
	FromName    string
	FromAddress string
	ToName      string
	ToAddress   string
}

type EmailTemplate struct {
	Title string
	Dogs  []*DogTemplate
}

type DogTemplate struct {
	Name  string
	Photo string
	Link  string
}

type Mailer interface {
	Send(subj string, content string) error
}

//go:embed templates/dogs.tmpl
var dogsTemplate []byte

func BuildHTML(dogs []*DogTemplate) (string, error) {
	tmpl, err := template.New("dogs").Parse(string(dogsTemplate))
	if err != nil {
		return "", errors.Wrap(err, "parsing template")
	}

	email := &EmailTemplate{
		Title: "New Dogs",
		Dogs:  dogs,
	}

	buf := &bytes.Buffer{}
	err = tmpl.Execute(buf, email)
	if err != nil {
		return "", errors.Wrap(err, "executing template")
	}
	return buf.String(), nil
}
