package main

import (
	"net/http"
	"projetoEmail/internal/domain/campaign"
	"projetoEmail/internal/endpoints"
	"projetoEmail/internal/infra/database"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Use(LoggerMiddleware)

	db := database.NewDB()

	service := campaign.Service{Repository: &database.CampaignRepository{Db: db}}
	handler := endpoints.Handler{CampaignService: &service}
	r.Route("/campaings", func(r chi.Router) {
		r.Post("/", handler.CampaignPost)
		r.Get("/", handler.CampaignsGet)
		r.Get("/{id}", handler.CampaignGet)
		r.Delete("/{id}", handler.CampaignDelete)
		r.Patch("/{id}/cancel", handler.CampaignCancelPatch)
	})
	http.ListenAndServe(":8080", r)
}

func LoggerMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		println(r.Method, r.URL.Path)

		next.ServeHTTP(w, r)
	})

}
