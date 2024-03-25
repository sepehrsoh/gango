package helloworld

import (
	"gango/utils"
)

type HelloWord struct {
	options utils.ServiceOptions
}

func NewHelloWord(configs ...utils.Options) HelloWord {
	opts := utils.DefaultOptions
	for _, config := range configs {
		config.Apply(&opts)
	}
	return HelloWord{options: opts}
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
	tmpl := utils.GetDefaultTemplateValues(name)
	if h.options.WithGrpc {
		tmpl["grpc"] = true
	} else {
		tmpl["gin"] = true
	}
	return tmpl
}
