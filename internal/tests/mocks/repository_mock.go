package mocks

import (
	"projetoEmail/internal/domain/campaign"

	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func (r *RepositoryMock) Save(campaign *campaign.Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func (r *RepositoryMock) All() ([]campaign.Campaign, error) {
	args := r.Called()

	return args.Get(0).([]campaign.Campaign), args.Error(1)
}

func (r *RepositoryMock) Get(id string) (campaign.Campaign, error) {
	args := r.Called(id)
	return args.Get(0).(campaign.Campaign), args.Error(1)
}

func (r *RepositoryMock) Update(id string, values *campaign.Campaign) error {
	args := r.Called(id, values)
	return args.Error(0)
}

func (r *RepositoryMock) Delete(id string) error {
	args := r.Called(id)
	return args.Error(0)
}
