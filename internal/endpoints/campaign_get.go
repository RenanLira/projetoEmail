package endpoints

import (
	"net/http"
	"projetoEmail/internal/utils"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) CampaignGet(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	campaign, err := h.CampaignService.Get(id)
	if err != nil {
		utils.SendJSON(w, err, nil)
		return
	}

	utils.SendJSON(w, nil, campaign)
}
