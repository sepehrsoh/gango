package wiring

import (
	"gango/utils"
)

type Wiring struct {
	options utils.ServiceOptions
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
	tmpl := utils.GetDefaultTemplateValues(name)
	if w.options.WithRedis {
		tmpl["withRedis"] = true
	}
	if w.options.WithElastic {
		tmpl["withElastic"] = true
	}
	if w.options.WithPostgres {
		tmpl["withPostgres"] = true
	}
	return tmpl
}

func NewWiring(configs ...utils.Options) Wiring {
	opts := utils.DefaultOptions
	for _, config := range configs {
		config.Apply(&opts)
	}
	return Wiring{options: opts}
}
