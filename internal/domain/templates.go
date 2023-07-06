package domain

import "io"

type TemplateEngine interface {
	Render(w io.Writer, templateName string, ctx interface{}) error
}
