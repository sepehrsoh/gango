package monitor

import (
	"gango/utils"
	"strings"
)

type Monitors struct{}

func (m Monitors) WriteFolder(dir string) error {
	return utils.EnrichTemplate(dir, m)
}

func (m Monitors) FileName() string {
	return "prometheus.go"
}

func (m Monitors) FilePath() string {
	return "/src/lib/monitor"
}

func (m Monitors) TemplateName() string {
	return "monitors.tmpl"
}

func (m Monitors) TemplateData(name string) map[string]interface{} {
	projectName := utils.GetProjectDir(name)
	namespace := strings.ToUpper(projectName[:1]) + projectName[1:]
	return map[string]interface{}{
		"ProjectName": name,
		"NameSpace":   namespace,
	}
}
