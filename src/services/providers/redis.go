package providers

import (
	"gango/utils"
)

type Redis struct {
}

func (r Redis) WriteFolder(dir string) error {
	return utils.EnrichTemplate(dir, r)
}

func (r Redis) FilePath() string {
	return "src/service/providers"
}

func (r Redis) FileName() string {
	return "redis.go"
}

func (r Redis) TemplateName() string {
	return "redisFile.tmpl"
}

func (r Redis) TemplateData(name string) map[string]interface{} {
	tmpl := utils.GetDefaultTemplateValues(name)
	return tmpl
}
