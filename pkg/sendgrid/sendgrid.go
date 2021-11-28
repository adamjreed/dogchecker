package sendgrid

import (
	"dogchecker/pkg/mailer"
	"fmt"
	"github.com/friendsofgo/errors"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"log"
)

type SendGridMailer struct {
	client *sendgrid.Client
	from   *mail.Email
	to     *mail.Email
}

func NewMailer(config *mailer.Config) *SendGridMailer {
	return &SendGridMailer{
		client: sendgrid.NewSendClient(config.ApiKey),
		from:   mail.NewEmail(config.FromName, config.FromAddress),
		to:     mail.NewEmail(config.ToName, config.ToAddress),
	}
}

func (s *SendGridMailer) Send(subject string, content string) error {
	message := mail.NewSingleEmail(s.from, subject, s.to, "", content)
	response, err := s.client.Send(message)
	if err != nil {
		return errors.Wrap(err, "calling sendgrid")
	}

	if response.StatusCode > 299 {
		return errors.New(fmt.Sprintf("invalid response (code: %d body: %s)", response.StatusCode, response.Body))
	}

	log.Println(fmt.Sprintf("email sent (code: %d body: %s)", response.StatusCode, response.Body))
	return nil
}
