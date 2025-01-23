package main

import (
	"net/http"
	"projetoEmail/internal/domain/campaign"
	"projetoEmail/internal/endpoints"
	"projetoEmail/internal/infra/database"
	"projetoEmail/internal/infra/mail"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {

	r := chi.NewRouter()

	r.Use(LoggerMiddleware)

	db := database.NewDB()

	service := campaign.Service{Repository: &database.CampaignRepository{Db: db}, SendMail: mail.SendMail}
	handler := endpoints.Handler{CampaignService: &service}
	r.Route("/campaings", func(r chi.Router) {
		authHandler := endpoints.NewAuthHandler()
		r.Use(authHandler.Auth)

		r.Post("/", handler.CampaignPost)
		r.Get("/", handler.CampaignsGet)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", handler.CampaignGet)
			r.Delete("/", handler.CampaignDelete)
			r.Patch("/cancel", handler.CampaignCancelPatch)
			r.Patch("/start", handler.CampaignStart)
		})
	})
	http.ListenAndServe(":8080", r)
}

func LoggerMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		println(r.Method, r.URL.Path)

		next.ServeHTTP(w, r)
	})

}
