package wiring

import (
	"gango/utils"
	"path/filepath"
	"strings"
)

type Wiring struct {
}

func (w Wiring) WriteFolder(dir string) error {
	return utils.WriteFile(dir, filepath.Join(w.FilePath(), w.FileName()), strings.ReplaceAll(wiringFile, "gango", dir))
}

func (w Wiring) FilePath() string {
	return "src/service/wiring"
}

func (w Wiring) FileName() string {
	return "wiring.go"
}

var wiringFile = `
package wiring

import "project/src/service/configs"


func NewWire() *Wire {
	wire := &Wire{
		config: configs.NewConfigFromEnv(),
	}
	wire.init()
	return wire
}

type Wire struct {
	config configs.Config
}

func (w *Wire) initProviders() {
	//	TODO implement providers registration
}

func (w *Wire) initRepositoryRegistry() {
	//	TODO implement command repository registration
}

func (w *Wire) init() {
	w.initProviders()
	w.initMetrics()
	w.initRepositoryRegistry()
}

func (w *Wire) initMetrics() {
	//	TODO implement metrics registration
}
`
