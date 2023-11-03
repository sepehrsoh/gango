package cmd

import (
	"gango/utils"
)

type Root struct {
}

func (r Root) WriteFolder(dir string) error {
	return utils.EnrichTemplate(dir, r)
}

func (r Root) FilePath() string {
	return "/cmd/base"
}

func (r Root) FileName() string {
	return "root.go"
}

func (r Root) TemplateName() string {
	return "rootFile.tmpl"
}

func (r Root) TemplateData(name string) map[string]interface{} {
	return map[string]interface{}{
		"ProjectName": name,
	}
}
