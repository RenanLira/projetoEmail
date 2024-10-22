package campaign

import (
	"testing"
	"time"

	"github.com/jaswdr/faker/v2"
	"github.com/stretchr/testify/assert"
)

var (
	fake    = faker.New()
	name    = "My Campaign"
	content = fake.Lorem().Text(100)
	emails  = []string{"tste@email.com", "email2@teste.com"}
)

func Test_New_CreateNew(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := New(name, content, emails)

	assert.Equal(campaign, &Campaign{
		ID:        campaign.ID,
		Name:      name,
		CreatedAt: campaign.CreatedAt,
		Content:   content,
		Contacts:  []Contact{{Email: "tste@email.com"}, {Email: "email2@teste.com"}},
	})
}

func Test_New_IDIsNotEmpty(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := New(name, content, emails)

	assert.NotEmpty(campaign.ID)
}

func Test_New_CreatedAtMustBeNow(t *testing.T) {
	assert := assert.New(t)

	now := time.Now().Add(-time.Second)

	campaign, _ := New(name, content, emails)

	assert.Greater(campaign.CreatedAt, now)
}

func Test_New_MustValidateName(t *testing.T) {
	assert := assert.New(t)

	_, err := New("", content, emails)

	assert.EqualError(err, "The field Name must have at least 5 characters")
}
