package misc

import (
	"gango/utils"
)

type Signals struct{}

func (s Signals) FileName() string {
	return "signals.go"
}

func (s Signals) WriteFolder(dir string) error {
	return utils.EnrichTemplate(dir, s)
}

func (s Signals) FilePath() string {
	return "/src/lib/misc"
}

func (s Signals) TemplateName() string {
	return "misc-signals.tmpl"
}

func (s Signals) TemplateData(name string) map[string]interface{} {
	return map[string]interface{}{
	}
}
