package wiring

import (
	"gango/utils"
)

type Internal struct {
	options utils.ServiceOptions
}

func NewInternal(configs ...utils.Options) Internal {
	opts := utils.DefaultOptions
	for _, config := range configs {
		config.Apply(&opts)
	}
	return Internal{options: opts}
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
	if i.options.WithGrpc {
		tmpl["grpc"] = true
	} else {
		tmpl["gin"] = true
	}
	return tmpl
}
