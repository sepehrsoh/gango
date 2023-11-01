package cmd

import (
	"gango/utils"
)

type ProjectName struct {
}

func (p ProjectName) WriteFolder(dir string) error {
	return utils.EnrichTemplate(dir, p)
}

func (p ProjectName) FilePath() string {
	return "/cmd"
}

func (p ProjectName) FileName() string {
	return "base.go"
}

func (p ProjectName) TemplateName() string {
	return "mainFile.tmpl"
}

func (p ProjectName) TemplateData(name string) map[string]interface{} {
	return map[string]interface{}{
		"ProjectName": name,
	}
}
