package endpoints

import "projetoEmail/internal/domain/campaign"

type Handler struct {
	CampaignService campaign.IService
}
