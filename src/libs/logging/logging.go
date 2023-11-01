package logging

import (
	"gango/utils"
)

type Logging struct {
}

func (l Logging) WriteFolder(dir string) error {
	return utils.EnrichTemplate(dir, l)
}

func (l Logging) FilePath() string {
	return "/src/lib/logging"
}

func (l Logging) FileName() string {
	return "logging.go"
}

func (l Logging) TemplateName() string {
	return "loggingFile.tmpl"
}

func (l Logging) TemplateData(name string) map[string]interface{} {
	return map[string]interface{}{
	}
}