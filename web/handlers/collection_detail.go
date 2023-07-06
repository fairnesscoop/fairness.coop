package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"fairness.coop/victor/internal/application"
	"fairness.coop/victor/internal/domain"
	"fairness.coop/victor/internal/domain/entities"
)

type CollectionDetailContext struct {
	Collection *entities.Collection
}

func RegisterHandleGetCollectionDetail(r *chi.Mux, ctn domain.Container) {
	r.Get("/collections/{collectionId}", func(w http.ResponseWriter, r *http.Request) {
		collectionId, err := strconv.Atoi(chi.URLParam(r, "collectionId"))

		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		collection, err := application.GetCollectionById(ctn, collectionId)

		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		err = ctn.GetTemplateEngine().Render(w, "collection_detail.html", CollectionDetailContext{collection})

		if err != nil {
			fmt.Println("error:", err.Error())
		}
	})
}
