package service

import (
	"fmt"
	"mime"
	"mime/multipart"
	"net/smtp"
	"path/filepath"
	"strings"

	"github.com/jordan-wright/email"
	"github.com/ymoldabe/Doodocs-Backend-Challenge/configs"
	"github.com/ymoldabe/Doodocs-Backend-Challenge/internal/models"
)

type MailType struct {
	cnf configs.Config
}

func NewMail(cnf configs.Config) *MailType {
	return &MailType{
		cnf: cnf,
	}
}

func (m *MailType) SendLetters(file *multipart.FileHeader, emailStr string) error {
	switch file.Header.Get("Content-Type") {
	case "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
		"application/pdf":
		emails := mailSeparator(emailStr)

		if err := m.sendMail(file, emails); err != nil {
			return err
		}
	default:
		return models.ErrBadRequest

	}
	return nil
}

func mailSeparator(emails string) []string {
	return strings.Split(emails, ",")
}

func (m *MailType) sendMail(file *multipart.FileHeader, emails []string) error {

	var (
		emailName     = m.cnf.EmailSenderName
		emailAddres   = m.cnf.EmailSenderAddres
		emailPassword = m.cnf.EmailSenderPassword
	)

	sender := NewGmailSender(emailName, emailAddres, emailPassword)

	subject := "IT IS VERY IMPORTANT!"
	content := `
	<h1>IT IS VERY IMPORTANT!</h1>
	<p>You can do this</p>`
	to := emails
	attachFile := file
	err := sender.SendEmail(subject, content, to, attachFile)
	return err
}

func (sender *GmailSender) SendEmail(
	subject string,
	content string,
	to []string,
	attachFiles *multipart.FileHeader,
) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", sender.name, sender.fromEmailAdders)
	e.Subject = subject
	e.HTML = []byte(content)
	e.To = to

	file, err := attachFiles.Open()
	if err != nil {
		return err

	}

	ct := mime.TypeByExtension(filepath.Ext(attachFiles.Filename))

	_, err = e.Attach(file, attachFiles.Filename, ct)
	if err != nil {
		return err
	}
	smtpAuth := smtp.PlainAuth("", sender.fromEmailAdders, sender.fromEmailPassword, stmpAuthAddres)
	return e.Send(stmpServerAddres, smtpAuth)

}

type EmailSen interface {
	SendEmail(
		subject string,
		content string,
		to []string,
		attachFiles *multipart.FileHeader,
	) error
}

type GmailSender struct {
	name              string
	fromEmailAdders   string
	fromEmailPassword string
}

func NewGmailSender(name, fromEmailAdders, fromEmailPassword string) EmailSen {
	return &GmailSender{
		name:              name,
		fromEmailAdders:   fromEmailAdders,
		fromEmailPassword: fromEmailPassword,
	}
}

const (
	stmpAuthAddres   = "smtp.gmail.com"
	stmpServerAddres = "smtp.gmail.com:587"
)
