package infrastructure

import (
	"fmt"
	"html/template"
	"io"
)

type FileTemplateEngine struct {
	homeTmpl             *template.Template
	collectionDetailTmpl *template.Template
	fileDetailTmpl       *template.Template
}

func NewFileTemplateEngine() (*FileTemplateEngine, error) {
	homeTmpl := template.Must(template.ParseFiles("web/handlers/home.html"))
	collectionDetailTmpl := template.Must(template.ParseFiles("web/handlers/collection_detail.html"))
	fileDetailTmpl := template.Must(template.ParseFiles("web/handlers/file_detail.html"))

	return &FileTemplateEngine{homeTmpl, collectionDetailTmpl, fileDetailTmpl}, nil
}

func (e *FileTemplateEngine) Render(w io.Writer, templateName string, ctx interface{}) error {
	switch templateName {
	case "home.html":
		return e.homeTmpl.Execute(w, ctx)
	case "collection_detail.html":
		return e.collectionDetailTmpl.Execute(w, ctx)
	case "file_detail.html":
		return e.fileDetailTmpl.Execute(w, ctx)
	default:
		return fmt.Errorf("template %s does not exist", templateName)
	}
}
