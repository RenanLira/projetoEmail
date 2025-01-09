package internalerrors

type ErrCampaignNotPending struct {
	CampaignID string
}

func (e ErrCampaignNotPending) Error() string {
	return "campaign " + e.CampaignID + " is not pending"
}

func NewErrCampaignNotPending(campaignID string) ErrCampaignNotPending {
	return ErrCampaignNotPending{CampaignID: campaignID}
}
