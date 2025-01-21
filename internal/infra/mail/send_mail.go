package mail

import (
	"os"
	"projetoEmail/internal/domain/campaign"

	"gopkg.in/gomail.v2"
)

func SendMail(campaign *campaign.Campaign) error {
	provider := gomail.NewDialer(os.Getenv("EMAIL_SMTP"), 587, os.Getenv("EMAIL_USER"), os.Getenv("EMAIL_PASSWORD"))

	emails := getEmailsFromContacts(campaign.Contacts)

	m := gomail.NewMessage()
	m.SetAddressHeader("From", campaign.CreatedBy, "Renan Lira")
	m.SetHeader("To", emails...)
	m.SetHeader("Subject", campaign.Name)
	m.SetBody("text/html", campaign.Content)

	return provider.DialAndSend(m)
}

func getEmailsFromContacts(contacts []campaign.Contact) []string {

	var emails []string

	for _, contact := range contacts {
		emails = append(emails, contact.Email)
	}

	return emails
}
