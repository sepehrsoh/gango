package registry

import (
	"os"
	"path/filepath"

	"gango/lib"

	"github.com/sirupsen/logrus"
)

var logger = lib.GetLogger("registry")

type Registry struct {
	Files map[string]lib.IWriteFolder
}

func NewRegistry() *Registry {
	return &Registry{Files: map[string]lib.IWriteFolder{}}
}

func (r *Registry) Register(name string, file lib.IWriteFolder) {
	r.Files[name] = file
}

func (r *Registry) Run(basePath string) {
	logger.Infoln("Start build project at new directory:", basePath)
	if basePath == "" {
		basePath = "./"
	}
	for name, file := range r.Files {
		logrus.Infoln("Run", name)
		dirPath := filepath.Join(basePath, file.FilePath())
		if err := os.MkdirAll(dirPath, 0750); err != nil {
			logrus.Error(err)
		}
		err := file.WriteFolder(basePath)
		if err != nil {
			logrus.Error(err)
		}
	}
}
