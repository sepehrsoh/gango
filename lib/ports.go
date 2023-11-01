package lib

import "go.uber.org/zap"

type LoggerFactory interface {
	New() *zap.SugaredLogger
}

type IWriteFolder interface {
	WriteFolder(dir string) error
	FilePath() string
}

type IWriteTemplate interface {
	FilePath() string
	FileName() string
	TemplateName() string 
	TemplateData(name string) map[string]interface{} 
}
