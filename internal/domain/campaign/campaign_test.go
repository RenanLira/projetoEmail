package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	name    = "My Campaign"
	content = "My Content"
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

	assert.EqualError(err, "name is required")
}
