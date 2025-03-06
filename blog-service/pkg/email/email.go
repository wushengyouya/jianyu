package email

import (
	"crypto/tls"
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
)

type Email struct {
	*SMTPInfo
}

type SMTPInfo struct {
	Host     string
	Port     int
	IsSSL    bool
	UserName string
	Password string
	From     string
}

func NewEmail(info *SMTPInfo) *Email {
	return &Email{SMTPInfo: info}
}

// 发送邮件
func (e *Email) SendMail(to []string, subject, body string) error {
	eEngine := email.NewEmail()
	eEngine.From = e.From
	eEngine.To = to
	eEngine.Subject = subject
	eEngine.Text = []byte(body)
	return eEngine.SendWithTLS(fmt.Sprint(e.Host, ":", e.Port), smtp.PlainAuth("", e.UserName, e.Password, e.Host),
		&tls.Config{InsecureSkipVerify: true, ServerName: e.Host},
	)
}
