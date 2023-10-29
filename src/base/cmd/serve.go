package cmd

import (
	"gango/utils"
)

type Serve struct{}

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
	return map[string]interface{}{
		"ProjectName": name,
	}
}