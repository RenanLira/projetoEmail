package campaign_test

import (
	"projetoEmail/internal/domain/campaign"
	"projetoEmail/internal/tests/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	sendMail = func(campaign *campaign.Campaign) error {
		return nil
	}
	repositoryMock *mocks.RepositoryMock
	service        campaign.Service
	campaignMock   *campaign.Campaign
)

func setUp() {
	repositoryMock = new(mocks.RepositoryMock)
	service = campaign.Service{
		Repository: repositoryMock,
		SendMail:   sendMail,
	}

	campaignMock, _ = campaign.New("My Campaign", "My Content", []string{"test@email.com"}, "creator@email.com")
}

func Test_CreateCampaign(t *testing.T) {
	setUp()

	dto := campaign.NewCampaignDTO{
		Name:      "My Campaign",
		Content:   "My Content",
		CreatedBy: "test@email.com",
		Emails:    []string{"teste@email.com"},
	}

	repositoryMock.On("Save", mock.MatchedBy(func(campaign *campaign.Campaign) bool {
		return true
	})).Return(nil)

	_, err := service.Create(dto)

	assert.Nil(t, err)
	repositoryMock.AssertExpectations(t)

}
