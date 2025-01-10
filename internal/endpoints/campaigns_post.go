package endpoints

import (
	"net/http"

	"projetoEmail/internal/domain/campaign"
	"projetoEmail/internal/utils"

	"github.com/go-chi/render"
)

func (h *Handler) CampaignPost(w http.ResponseWriter, r *http.Request) {
	var request campaign.NewCampaignDTO

	render.DecodeJSON(r.Body, &request)

	id, err := h.CampaignService.Create(request)

	if err != nil {
		utils.SendJSON(w, err, nil)
		return
	}

	res := &utils.Success{Status: http.StatusCreated, Data: id}

	utils.SendJSON(w, nil, res)
}
