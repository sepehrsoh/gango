package wiring

import (
	"gango/utils"
)

type Service struct {
	options utils.ServiceOptions
}

func NewService(configs ...utils.Options) Service {
	opts := utils.DefaultOptions
	for _, config := range configs {
		config.Apply(&opts)
	}
	return Service{options: opts}
}

func (s Service) WriteFolder(dir string) error {
	return utils.EnrichTemplate(dir, s)
}

func (s Service) FilePath() string {
	return "src/service/wiring"
}

func (s Service) FileName() string {
	return "service.go"
}

func (s Service) TemplateName() string {
	return "serviceFile.tmpl"
}

func (s Service) TemplateData(name string) map[string]interface{} {
	tmpl := utils.GetDefaultTemplateValues(name)
	if s.options.WithRedis {
		tmpl["withRedis"] = true
	}
	if s.options.WithElastic {
		tmpl["withElastic"] = true
	}
	if s.options.WithPostgres {
		tmpl["withPostgres"] = true
	}
	if s.options.WithGrpc {
		tmpl["grpc"] = true
	} else {
		tmpl["gin"] = true
	}
	return tmpl
}
