package executors

import (
	"gango/utils"
	"path/filepath"
	"strings"
)

type Registry struct {
}

func (e Registry) WriteFolder(dir string) error {
	return utils.WriteFile(dir, filepath.Join(e.FilePath(), e.FileName()), strings.ReplaceAll(registryFile, "gango", dir))
}

func (e Registry) FilePath() string {
	return "/src/lib/executors"
}

func (e Registry) FileName() string {
	return "registry.go"
}

var registryFile = `
package executors

import (
	"os"
	"os/signal"
	"syscall"

	"gango/src/lib/logging"
)

var logger = logging.GetLogger("executors")

type IWorker interface {
	Loop() error
}

type IExecutor interface {
	Start(interrupt <-chan os.Signal)
}


type Registry struct {
	executors []IExecutor
}

func (w *Registry) Register(executor IExecutor) *Registry {
	w.executors = append(w.executors, executor)
	return w
}

func (w *Registry) Start() {
	logger.Info("Starting workers...")
	for _, executor := range w.executors {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

		go executor.Start(c)
	}
}

`
