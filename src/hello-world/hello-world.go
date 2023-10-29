package hello_world

import (
	"gango/utils"
)

type HelloWord struct {
}

func (h HelloWord) WriteFolder(dir string) error {
	return utils.EnrichTemplate(dir, h)
}

func (h HelloWord) FilePath() string {
	return "/src/hello-world"
}

func (h HelloWord) FileName() string {
	return "hello-world.go"
}

func (h HelloWord) TemplateName() string {
	return "helloWorldFile.tmpl"
}

func (h HelloWord) TemplateData(name string) map[string]interface{} {
	return map[string]interface{}{
		"ProjectName": name,
	}
}
