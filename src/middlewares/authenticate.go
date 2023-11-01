package middlewares

import (
	"gango/utils"
)

type Middleware struct{}

func (m Middleware) FileName() string {
	return "middleware.go"
}

func (m Middleware) WriteFolder(dir string) error {
	return utils.EnrichTemplate(dir, m)
}

func (m Middleware) FilePath() string {
	return "/src/middlewares"
}

func (m Middleware) TemplateName() string {
	return "middleware.tmpl"
}

func (m Middleware) TemplateData(name string) map[string]interface{} {
	return map[string]interface{}{
	}
}
