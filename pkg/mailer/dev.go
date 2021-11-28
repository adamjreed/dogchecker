package mailer

import (
	"fmt"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"log"
)

type DevMailer struct {
	error error
	from  *mail.Email
	to    *mail.Email
}

func NewDevMailer(config *Config) *DevMailer {
	return &DevMailer{
		from: mail.NewEmail(config.FromName, config.FromAddress),
		to:   mail.NewEmail(config.ToName, config.ToAddress),
	}
}

func (d *DevMailer) Send(subject string, content string) error {
	if d.error != nil {
		return d.error
	}

	log.Println(fmt.Sprintf("pretending to send email (to %s (%s), from %s (%s), subject '%s', content '%s'", d.to.Name, d.to.Address, d.from.Name, d.from.Address, subject, content))
	return nil
}
