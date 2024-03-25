package cmd

import (
	"gango/utils"
)

type Serve struct {
	options utils.ServiceOptions
}

func NewServe(configs ...utils.Options) Serve {
	opts := utils.DefaultOptions
	for _, config := range configs {
		config.Apply(&opts)
	}
	return Serve{options: opts}
}

func (s Serve) WriteFolder(dir string) error {
	return utils.EnrichTemplate(dir, s)
}

func (s Serve) FilePath() string {
	return "/cmd/base"
}

func (s Serve) FileName() string {
	return "serve.go"
}

func (s Serve) TemplateName() string {
	return "serveFile.tmpl"
}

func (s Serve) TemplateData(name string) map[string]interface{} {
	tmpl := utils.GetDefaultTemplateValues(name)
	if s.options.WithGrpc {
		tmpl["grpc"] = true
	} else {
		tmpl["gin"] = true
	}
	return tmpl
}
