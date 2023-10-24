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

import (
	"time"

	"gango/src/lib/executors"
	"gango/src/service/configs"

	"github.com/go-redis/redis"
)


func NewWire() *Wire {
	wire := &Wire{
		config: configs.NewConfigFromEnv(),
	}
	wire.init()
	return wire
}

type Wire struct {
	config 				configs.Config
	executorsRegistry   *executors.Registry
	// add more providers like redis implementation here
	redis               *redis.Client
}

//	initProviders initials providers you are using
func (w *Wire) initProviders() {
	//  w.GetRedis()
}

func (w *Wire) init() {
	w.initProviders()
	w.initMetrics()
	w.initExecutorsRegistry()
}

func (w *Wire) initMetrics() {
	StartTime.Add(float64(time.Now().Unix()))
	metricsServer := w.GetMetricsServer()
	metricsServer.Register(executors.WorkerCounter)
	metricsServer.Register(StartTime)
}

func (w *Wire) initExecutorsRegistry() {
	registry := w.GetExecutorsRegistry()
	registry.Register(w.GetMetricsServer())
}
`
