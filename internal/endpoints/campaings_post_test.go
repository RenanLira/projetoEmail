package endpoints

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"projetoEmail/internal/domain/campaign"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ServiceMock struct {
	mock.Mock
}

func (s *ServiceMock) Create(dto campaign.NewCampaignDTO) (id string, err error) {
	args := s.Called(dto)

	return args.String(0), args.Error(1)
}
func (s *ServiceMock) All() ([]campaign.Campaign, error) {
	args := s.Called()

	return args.Get(0).([]campaign.Campaign), args.Error(1)
}

func (s *ServiceMock) Get(id string) (*campaign.GetCampaignDTO, error) {
	args := s.Called(id)

	return args.Get(0).(*campaign.GetCampaignDTO), args.Error(1)
}

func (s *ServiceMock) Cancel(id string) error {
	args := s.Called(id)

	return args.Error(0)
}

func (s *ServiceMock) Delete(id string) error {
	args := s.Called(id)

	return args.Error(0)
}

func Test_CampaignsPost_should_save_new_campaign(t *testing.T) {
	assert := assert.New(t)
	service := new(ServiceMock)
	handler := Handler{CampaignService: service}
	body := campaign.NewCampaignDTO{
		Name:    "My Campaign",
		Content: "My Content",
		Emails:  []string{"teste@email.com"},
	}

	service.On("Create", body).Return("123", nil)

	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)

	req, _ := http.NewRequest("POST", "/", &buf)
	res := httptest.NewRecorder()

	handler.CampaignPost(res, req)

	assert.Equal(http.StatusOK, res.Code)
}
