package campaign

import (
	"projetoEmail/internal/contract"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func Test_CreateCampaign(t *testing.T) {
	assert := assert.New(t)
	dto := contract.NewCampaignDTO{
		Name:    "My Campaign",
		Content: "My Content",
		Emails:  []string{"teste@email.com"},
	}

	repository := new(repositoryMock)
	repository.On("Save", mock.Anything).Return(nil)

	service := Service{Repository: repository}

	_, err := service.Create(dto)

	assert.Nil(err)
	repository.AssertExpectations(t)

}
