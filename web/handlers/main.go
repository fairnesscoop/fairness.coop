package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"fairness.coop/victor/internal/domain"
)

func NewRouter(ctn domain.Container) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RedirectSlashes)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	})

	RegisterHandleHome(r, ctn)
	RegisterHandleGetCollectionDetail(r, ctn)
	RegisterHandleGetFile(r, ctn)
	RegisterHandleUpdateFile(r, ctn)

	return r
}
