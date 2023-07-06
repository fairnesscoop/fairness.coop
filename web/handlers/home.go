package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	"fairness.coop/victor/internal/application"
	"fairness.coop/victor/internal/domain"
	"fairness.coop/victor/internal/domain/entities"
)

type adminContext struct {
	Collections []*entities.Collection
}

func RegisterHandleHome(r *chi.Mux, ctn domain.Container) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		collections := application.GetAllCollections(ctn)

		err := ctn.GetTemplateEngine().Render(w, "home.html", adminContext{collections})

		if err != nil {
			fmt.Println("error:", err.Error())
		}
	})
}
