package providers

import "gango/utils"

type Postgres struct {
}

func (r Postgres) WriteFolder(dir string) error {
	return utils.EnrichTemplate(dir, r)
}

func (r Postgres) FilePath() string {
	return "src/service/providers"
}

func (r Postgres) FileName() string {
	return "postgres.go"
}

func (r Postgres) TemplateName() string {
	return "postgresFile.tmpl"
}

func (r Postgres) TemplateData(name string) map[string]interface{} {
	tmpl := utils.GetDefaultTemplateValues(name)
	return tmpl
}
