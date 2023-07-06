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

type FileDetailContext struct {
	Collection *entities.Collection
	File       *entities.File
	Content    string
}

func RegisterHandleGetFile(r *chi.Mux, ctn domain.Container) {
	r.Get("/collections/{collectionId}/posts/*", func(w http.ResponseWriter, r *http.Request) {
		collectionId, err := strconv.Atoi(chi.URLParam(r, "collectionId"))

		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		path := chi.URLParam(r, "*")
		file, err := application.GetFileByPathAndCollectionId(ctn, path, collectionId)

		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		content, err := file.Read()
		if err != nil {
			fmt.Println("error:", err.Error())
			http.Error(w, http.StatusText(500), 500)
			return
		}

		err = ctn.GetTemplateEngine().Render(w, "file_detail.html", FileDetailContext{File: file, Content: string(content)})

		if err != nil {
			fmt.Println("error:", err)
		}
	})
}

func RegisterHandleUpdateFile(r *chi.Mux, ctn domain.Container) {
	r.Post("/collections/{collectionId}/posts/*", func(w http.ResponseWriter, r *http.Request) {
		collectionId, err := strconv.Atoi(chi.URLParam(r, "collectionId"))

		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		path := chi.URLParam(r, "*")
		file, err := application.GetFileByPathAndCollectionId(ctn, path, collectionId)

		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		err = r.ParseForm()

		if err != nil {
			http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
			return
		}

		content := r.PostForm.Get("content")
		err = file.Write(content)

		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, file.Url(), http.StatusSeeOther)
	})
}
