package endpoints

import (
	"net/http"
	"projetoEmail/internal/utils"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) CampaignStart(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	
	err := h.CampaignService.Start(id)
	if err != nil {
		utils.SendJSON(w, err, nil)
		return
	}

	res := &utils.Success{Status: http.StatusOK, Data: "Campaign started"}

	utils.SendJSON(w, nil, res)
}
