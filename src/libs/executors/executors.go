package executors

import (
	"gango/utils"
	"path/filepath"
	"strings"
)

type Execute struct {
}

func (e Execute) WriteFolder(dir string) error {
	return utils.WriteFile(dir, filepath.Join(e.FilePath(), e.FileName()), strings.ReplaceAll(executorFile, "gango", dir))
}

func (e Execute) FilePath() string {
	return "/src/lib/executors"
}

func (e Execute) FileName() string {
	return "executors.go"
}

var executorFile = `
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
