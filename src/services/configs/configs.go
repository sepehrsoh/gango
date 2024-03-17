package configs

import (
	"gango/utils"
)

type Config struct {
	options utils.ServiceOptions
}

func NewConfig(configs ...utils.Options) Config {
	opts := utils.DefaultOptions
	for _, config := range configs {
		config.Apply(&opts)
	}
	return Config{options: opts}
}

func (c Config) WriteFolder(dir string) error {
	return utils.EnrichTemplate(dir, c)
}

func (c Config) FilePath() string {
	return "src/service/configs"
}

func (c Config) FileName() string {
	return "configs.go"
}

func (c Config) TemplateName() string {
	return "configsFile.tmpl"
}

func (c Config) TemplateData(name string) map[string]interface{} {
	tmpl := utils.GetDefaultTemplateValues(name)
	if c.options.WithRedis {
		tmpl["withRedis"] = true
	}
	if c.options.WithElastic {
		tmpl["withElastic"] = true
	}
	if c.options.WithPostgres {
		tmpl["withPostgres"] = true
	}
	return tmpl
}
