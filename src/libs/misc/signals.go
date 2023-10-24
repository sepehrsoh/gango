package misc

import (
	"path/filepath"

	"gango/utils"
)

type Signals struct{}

func (s Signals) FileName() string {
	return "signals.go"
}

func (s Signals) WriteFolder(dir string) error {
	return utils.WriteFile(dir, filepath.Join(s.FilePath(), s.FileName()), signals)
}

func (s Signals) FilePath() string {
	return "/src/lib/misc"
}

var signals = `
package misc

import (
	"os"
	"os/signal"
	"syscall"
)

func CreateTerminateChannel() chan os.Signal {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	return c
}
`
