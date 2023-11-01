package wiring

import (
	"gango/utils"
)

type Wiring struct {
}

func (w Wiring) WriteFolder(dir string) error {
	return utils.EnrichTemplate(dir, w)
}

func (w Wiring) FilePath() string {
	return "src/service/wiring"
}

func (w Wiring) FileName() string {
	return "wiring.go"
}

func (w Wiring) TemplateName() string {
	return "wiringFile.tmpl"
}

func (w Wiring) TemplateData(name string) map[string]interface{} {
	return map[string]interface{}{
		"ProjectName": name,
	}
}
