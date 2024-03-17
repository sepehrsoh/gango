package providers

import "gango/utils"

type Elastic struct {
}

func (r Elastic) WriteFolder(dir string) error {
	return utils.EnrichTemplate(dir, r)
}

func (r Elastic) FilePath() string {
	return "src/service/providers"
}

func (r Elastic) FileName() string {
	return "elastic.go"
}

func (r Elastic) TemplateName() string {
	return "elasticFile.tmpl"
}

func (r Elastic) TemplateData(name string) map[string]interface{} {
	tmpl := utils.GetDefaultTemplateValues(name)
	return tmpl
}
