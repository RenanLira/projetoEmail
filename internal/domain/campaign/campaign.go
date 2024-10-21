package campaign

import (
	"time"

	"github.com/rs/xid"
)

type Contact struct {
	Email string `validate:"required,email"`
}

type Campaign struct {
	ID        string    `validate:"required"`
	Name      string    `validate:"min=5,max=100"`
	CreatedAt time.Time `validate:"required"`
	Content   string    `validate:"min=5"`
	Contacts  []Contact `validate:"min=1"`
}

func New(name string, content string, emails []string) (*Campaign, error) {

	contacts := make([]Contact, len(emails))
	for i, email := range emails {
		contacts[i] = Contact{Email: email}
	}

	return &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		CreatedAt: time.Now(),
		Content:   content,
		Contacts:  contacts,
	}, nil
}
