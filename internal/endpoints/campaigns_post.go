package endpoints

import (
	"fmt"
	"net/http"

	"projetoEmail/internal/domain/campaign"
	"projetoEmail/internal/utils"

	"github.com/go-chi/render"
)

func (h *Handler) CampaignPost(w http.ResponseWriter, r *http.Request) {
	var request campaign.NewCampaignDTO

	render.DecodeJSON(r.Body, &request)

	// Get email from context for CreatedBy field
	email := r.Context().Value("email").(string)
	fmt.Println("Email from context: ", email)
	request.CreatedBy = email

	id, err := h.CampaignService.Create(request)

	if err != nil {
		utils.SendJSON(w, err, nil)
		return
	}

	res := &utils.Success{Status: http.StatusCreated, Data: id}

	utils.SendJSON(w, nil, res)
}
