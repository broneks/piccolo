package rendererservice

import (
	"html/template"
)

type RendererService struct {
	templates *template.Template
}

func New(dirPath string) *RendererService {
	return &RendererService{
		templates: template.Must(template.ParseGlob(dirPath)),
	}
}
