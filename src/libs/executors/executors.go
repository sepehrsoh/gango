package executors

import (
	"gango/utils"
)

type Executors struct {
}

func (e Executors) WriteFolder(dir string) error {
	return utils.EnrichTemplate(dir, e)
}

func (e Executors) FilePath() string {
	return "/src/lib/executors"
}

func (e Executors) FileName() string {
	return "executors.go"
}

func (e Executors) TemplateName() string {
	return "executorFile.tmpl"
}

func (e Executors) TemplateData(name string) map[string]interface{} {
	return map[string]interface{}{
	}
}