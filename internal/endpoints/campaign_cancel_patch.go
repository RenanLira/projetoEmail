package endpoints

import (
	"net/http"
	"projetoEmail/internal/utils"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) CampaignCancelPatch(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := h.CampaignService.Cancel(id)
	if err != nil {
		utils.SendJSON(w, err, nil)
		return
	}

	utils.SendJSON(w, nil, map[string]string{"message": "Campaign cancelled"})
}
