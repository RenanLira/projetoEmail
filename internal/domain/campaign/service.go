package campaign

import (
	"projetoEmail/internal/contract"
	internalerrors "projetoEmail/internal/internal_errors"
)

type Service struct {
	Repository Repository
}

func (s *Service) Create(dto contract.NewCampaignDTO) (id string, err error) {

	campaign, err := New(dto.Name, dto.Content, dto.Emails)
	if err != nil {
		return
	}

	err = s.Repository.Save(campaign)
	if err != nil {
		err = internalerrors.ErrInternal
		return
	}

	id = campaign.ID

	return
}
