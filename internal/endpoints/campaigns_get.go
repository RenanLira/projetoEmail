package endpoints

import (
	"net/http"
	"projetoEmail/internal/utils"
)

func (h *Handler) CampaignsGet(w http.ResponseWriter, r *http.Request) {
	campaigns, err := h.CampaignService.All()

	if err != nil {
		utils.SendJSON(w, err, nil)
		return
	}

	utils.SendJSON(w, nil, campaigns)
}
