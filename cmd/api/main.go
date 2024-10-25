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

	service := campaign.Service{Repository: &database.CampaignRepository{}}
	handler := endpoints.Handler{CampaignService: service}

	r.Post("/campaings", handler.CampaignPost)

	http.ListenAndServe(":8080", r)
}

func LoggerMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		println(r.Method, r.URL.Path)

		next.ServeHTTP(w, r)
	})

}
