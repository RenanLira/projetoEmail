package internalerrors

import "net/http"

type ErrCampaignNotPending struct {
	HttpError
	CampaignID string
}

func (e ErrCampaignNotPending) Error() string {
	return "campaign " + e.CampaignID + " is not pending"
}

func NewErrCampaignNotPending(campaignID string) ErrCampaignNotPending {
	return ErrCampaignNotPending{CampaignID: campaignID, HttpError: HttpError{HttpStatus: http.StatusConflict}}
}
