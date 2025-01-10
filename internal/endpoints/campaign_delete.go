package endpoints

import (
	"net/http"
	"projetoEmail/internal/utils"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) CampaignDelete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := h.CampaignService.Delete(id)
	if err != nil {
		utils.SendJSON(w, err, nil)
		return
	}

	res := utils.Success{Status: http.StatusOK, Data: map[string]string{"message": "Campanha deletada com sucesso"}}

	utils.SendJSON(w, nil, &res)
}
