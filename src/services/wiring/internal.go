package wiring

import (
	"gango/utils"
)

type Internal struct {
}

func (i Internal) WriteFolder(dir string) error {
	return utils.EnrichTemplate(dir, i)
}

func (i Internal) FilePath() string {
	return "src/service/wiring"
}

func (i Internal) FileName() string {
	return "internal.go"
}

func (i Internal) TemplateName() string {
	return "internalFile.tmpl"
}

func (i Internal) TemplateData(name string) map[string]interface{} {
	tmpl := utils.GetDefaultTemplateValues(name)
	return tmpl
}
