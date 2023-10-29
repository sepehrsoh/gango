package wiring

import (
	"gango/utils"
)

type Service struct {
}

func (s Service) WriteFolder(dir string) error {
	return utils.EnrichTemplate(dir, s)
}

func (s Service) FilePath() string {
	return "src/service/wiring"
}

func (s Service) FileName() string {
	return "service.go"
}

func (s Service) TemplateName() string {
	return "serviceFile.tmpl"
}

func (s Service) TemplateData(name string) map[string]interface{} {
	return map[string]interface{}{
		"ProjectName": name,
	}
}
