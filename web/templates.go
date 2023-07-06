package web

import "html/template"

type Templates struct {
	AdminTmpl            *template.Template
	CollectionDetailTmpl *template.Template
	FileDetailTmpl       *template.Template
}

func NewTemplates() (*Templates, error) {

	adminTmpl, err := template.ParseFiles("web/index.html")

	if err != nil {
		return nil, err
	}

	collectionDetailTmpl, err := template.ParseFiles("web/collection_detail.html")

	if err != nil {
		return nil, err
	}

	fileDetailTmpl, err := template.ParseFiles("web/file_detail.html")

	if err != nil {
		return nil, err
	}

	return &Templates{adminTmpl, collectionDetailTmpl, fileDetailTmpl}, nil
}
