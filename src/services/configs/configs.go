package configs

import (
	"gango/utils"
)

type Config struct{}

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
	return map[string]interface{}{}
}
