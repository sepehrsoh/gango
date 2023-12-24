package wiring

import (
	"gango/utils"
)

type Metrics struct {
}

func (m Metrics) WriteFolder(dir string) error {
	return utils.EnrichTemplate(dir, m)
}

func (m Metrics) FilePath() string {
	return "src/service/wiring"
}

func (m Metrics) FileName() string {
	return "metrics.go"
}

func (m Metrics) TemplateName() string {
	return "metricsFile.tmpl"
}

func (m Metrics) TemplateData(name string) map[string]interface{} {
	return map[string]interface{}{
		"ProjectName": name,
		"MetricName":  utils.GetProjectDir(name),
	}
}
