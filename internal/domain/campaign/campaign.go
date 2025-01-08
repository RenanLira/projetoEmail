package campaign

import (
	"fmt"
	internalerrors "projetoEmail/internal/internal_errors"
	"time"

	"github.com/rs/xid"
)

type Contact struct {
	ID         string `gorm:"size:50" json:"-"`
	Email      string `validate:"required,email" json:"email"`
	CampaignId string `json:"-"`
}

type Status uint8

func (s Status) String() string {
	statusString := [5]string{"Pending", "Cancelled", "Deleted", "Started", "Finished"}

	return statusString[s]
}

func (s Status) MarshalJSON() ([]byte, error) {
	return []byte(`"` + s.String() + `"`), nil
}

const (
	Pending Status = iota
	Cancelled
	Deleted
	Started
	Finished
)

type Campaign struct {
	ID        string    `validate:"required" json:"id"`
	Name      string    `validate:"min=5,max=100" gorm:"size:100" json:"name"`
	CreatedAt time.Time `validate:"required" json:"created_at"`
	Content   string    `validate:"min=5" gorm:"size:500" json:"content"`
	Contacts  []Contact `validate:"min=1,dive" json:"contacts,omitempty" gorm:"OnDelete:CASCADE"`
	Status    Status    `gorm:"" json:"status"`
}

func (c *Campaign) Cancel() {
	c.Status = Cancelled
}

func (c *Campaign) Delete() {
	c.Status = Deleted
}

func New(name string, content string, emails []string) (campaign *Campaign, err error) {

	contacts := make([]Contact, len(emails))
	for i, email := range emails {
		contacts[i] = Contact{Email: email, ID: xid.New().String()}
	}

	campaign = &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		CreatedAt: time.Now(),
		Content:   content,
		Contacts:  contacts,
		Status:    Pending,
	}

	fmt.Println(campaign)

	err = internalerrors.ValidateStruct(campaign)

	if err != nil {
		return nil, err
	}

	return campaign, err
}
