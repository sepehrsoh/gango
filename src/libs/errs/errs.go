package errs

import (
	"gango/utils"
)

type Errors struct{}

func (s Errors) WriteFolder(dir string) error {
	return utils.EnrichTemplate(dir, s)
}

func (s Errors) FileName() string {
	return "errs.go"
}

func (s Errors) FilePath() string {
	return "/src/lib/errs"
}


func (s Errors) TemplateName() string {
	return "signals.tmpl"
}

func (s Errors) TemplateData(name string) map[string]interface{} {
	return map[string]interface{}{
		"ProjectName": name,
	}
}
