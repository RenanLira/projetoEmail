package campaign

import (
	"testing"
	"time"

	"github.com/jaswdr/faker/v2"
	"github.com/stretchr/testify/assert"
)

var (
	fake       = faker.New()
	name       = "My Campaign"
	content    = fake.Lorem().Text(100)
	emails     = []string{"tste@email.com", "email2@teste.com"}
	created_by = "criador@teste.com"
)

func Test_New_CreateNew(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := New(name, content, emails, created_by)

	assert.Equal(campaign, &Campaign{
		ID:        campaign.ID,
		Name:      name,
		CreatedAt: campaign.CreatedAt,
		Content:   content,
		CreatedBy: created_by,
		Contacts:  []Contact{{Email: "tste@email.com", ID: campaign.Contacts[0].ID, CampaignId: campaign.ID}, {Email: "email2@teste.com", ID: campaign.Contacts[1].ID, CampaignId: campaign.ID}},
	})
}

func Test_New_IDIsNotEmpty(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := New(name, content, emails, created_by)

	assert.NotEmpty(campaign.ID)
}

func Test_New_CreatedAtMustBeNow(t *testing.T) {
	assert := assert.New(t)

	now := time.Now().Add(-time.Second)

	campaign, _ := New(name, content, emails, created_by)

	assert.Greater(campaign.CreatedAt, now)
}

func Test_New_MustValidateName(t *testing.T) {
	assert := assert.New(t)

	_, err := New("", content, emails, created_by)

	assert.EqualError(err, "The field Name must have at least 5 characters")
}
