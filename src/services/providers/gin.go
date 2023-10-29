package providers

import (
	"gango/utils"
)

type Gin struct {
}

func (g Gin) WriteFolder(dir string) error {
	return utils.EnrichTemplate(dir, g)
}

func (g Gin) FilePath() string {
	return "src/service/providers"
}

func (g Gin) FileName() string {
	return "gin.go"
}

func (g Gin) TemplateName() string {
	return "ginFile.tmpl"
}

func (g Gin) TemplateData(name string) map[string]interface{} {
	return map[string]interface{}{}
}
