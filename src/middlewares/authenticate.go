package middlewares

import (
	"gango/utils"
)

type Middleware struct {
	options utils.ServiceOptions
}

func NewMiddleware(configs ...utils.Options) Middleware {
	opts := utils.DefaultOptions
	for _, config := range configs {
		config.Apply(&opts)
	}
	return Middleware{options: opts}
}

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
	tmpl := utils.GetDefaultTemplateValues(name)
	if m.options.WithGrpc {
		tmpl["grpc"] = true
	} else {
		tmpl["gin"] = true
	}
	return tmpl
}
