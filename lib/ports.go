package lib

import "go.uber.org/zap"

type LoggerFactory interface {
	New() *zap.SugaredLogger
}

type IWriteFolder interface {
	WriteFolder(dir string) error
	FilePath() string
	FileName() string
}
