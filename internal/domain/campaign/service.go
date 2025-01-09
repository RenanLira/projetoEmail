package campaign

import (
	internalerrors "projetoEmail/internal/internal_errors"
)

type IService interface {
	Create(dto NewCampaignDTO) (id string, err error)
	All() ([]Campaign, error)
	Get(id string) (*GetCampaignDTO, error)
	Cancel(id string) error
	Delete(id string) error
}

type Service struct {
	Repository Repository
}

func (s *Service) Create(dto NewCampaignDTO) (id string, err error) {

	campaign, err := New(dto.Name, dto.Content, dto.Emails)
	if err != nil {
		return
	}

	err = s.Repository.Save(campaign)
	if err != nil {
		err = internalerrors.NewErrInternal(err.Error())
		return
	}

	id = campaign.ID

	return
}

func (s *Service) All() ([]Campaign, error) {
	return s.Repository.All()
}

func (s *Service) Get(id string) (*GetCampaignDTO, error) {
	campaign, err := s.Repository.Get(id)
	if err != nil {
		return nil, err
	}

	return &GetCampaignDTO{
		ID:      campaign.ID,
		Send:    campaign.Contacts,
		Content: campaign.Content,
		Status:  campaign.Status.String(),
		Amount:  len(campaign.Contacts),
	}, nil
}

func (s *Service) Cancel(id string) error {

	campaign, err := s.Repository.Get(id)
	if err != nil {
		return err
	}

	if campaign.Status != Pending {
		return internalerrors.NewErrCampaignNotPending(id)
	}

	campaign.Cancel()
	err = s.Repository.Update(id, &campaign)
	if err != nil {
		return internalerrors.NewErrInternal(err.Error())
	}

	return nil
}

func (s *Service) Delete(id string) error {

	campaign, err := s.Repository.Get(id)
	if err != nil {
		return err
	}

	campaign.Delete()
	err = s.Repository.Update(id, &campaign)
	if err != nil {
		return internalerrors.NewErrInternal(err.Error())
	}

	return nil
}
