package executors

import (
	"gango/utils"
)

type Registry struct {
}

func (e Registry) WriteFolder(dir string) error {
	return utils.EnrichTemplate(dir, e)
}

func (e Registry) FilePath() string {
	return "/src/lib/executors"
}

func (e Registry) FileName() string {
	return "registry.go"
}

func (e Registry) TemplateName() string {
	return "registryFile.tmpl"
}

func (e Registry) TemplateData(name string) map[string]interface{} {
	return map[string]interface{}{
		"ProjectName": name,
	}
}
