package campaign

import (
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

func (r *repositoryMock) All() ([]Campaign, error) {
	args := r.Called()
	return args.Get(0).([]Campaign), args.Error(1)
}

func (r *repositoryMock) Get(id string) (Campaign, error) {
	args := r.Called(id)
	return args.Get(0).(Campaign), args.Error(1)
}

func (r *repositoryMock) Update(id string, values *Campaign) error {
	args := r.Called(id, values)
	return args.Error(0)
}

func (r *repositoryMock) Delete(id string) error {
	args := r.Called(id)
	return args.Error(0)
}

func Test_CreateCampaign(t *testing.T) {
	assert := assert.New(t)
	dto := NewCampaignDTO{
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
