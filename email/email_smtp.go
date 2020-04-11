package email

import (
	"errors"
	"net/smtp"

	"github.com/jordan-wright/email"
)

// SMTPConfig - config parameters
type SMTPConfig struct {
	SMTPHost     string
	SMTPPort     string
	SMTPUsername string
	SMTPPassword string
	SMTPFrom     string
}

// SendEmail - sends emails
func SendEmail(config *SMTPConfig, toEmail string, subject string, body string) error {

	//Return if email configuration not complete
	if len(config.SMTPUsername) == 0 || len(config.SMTPPassword) == 0 || len(config.SMTPHost) == 0 || len(config.SMTPPort) == 0 {
		return errors.New("SMTP Host, Port, Username or Password empty")
	}

	e := email.NewEmail()
	e.From = config.SMTPFrom
	e.To = []string{toEmail}
	e.Subject = subject
	e.HTML = []byte(body)
	return e.Send(config.SMTPHost+":"+config.SMTPPort, smtp.PlainAuth("", config.SMTPUsername, config.SMTPPassword, config.SMTPHost))
}
