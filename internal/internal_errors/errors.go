package internalerrors

import "errors"

var ErrInternal error = errors.New("internal server error")

var ErrCampaignNotPending error = errors.New("campaign is not pending")

var ErrCampaignNotFound error = errors.New("campaign not found")
