package wiring

import (
	"gango/utils"
	"path/filepath"
	"strings"
)

type Internal struct {
}

func (i Internal) WriteFolder(dir string) error {
	return utils.WriteFile(dir, filepath.Join(i.FilePath(), i.FileName()), strings.ReplaceAll(internalFile, "gango", dir))
}

func (i Internal) FilePath() string {
	return "src/service/wiring"
}

func (i Internal) FileName() string {
	return "internal.go"
}

var internalFile = `
package wiring

import "gango/src/lib/monitor"

func (w *Wire) GetMetricsServer() *monitor.PrometheusMetricsServer {
	return &monitor.PrometheusMetricsServer{
		Url:  w.config.Monitor.Host,
		Port: w.config.Monitor.Port,
	}
}

`
