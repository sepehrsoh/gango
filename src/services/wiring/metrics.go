package wiring

import (
	"gango/utils"
	"path/filepath"
	"strings"
)

type Metrics struct {
}

func (m Metrics) WriteFolder(dir string) error {
	return utils.WriteFile(dir, filepath.Join(m.FilePath(), m.FileName()), strings.ReplaceAll(metricsFile, "gango", dir))
}

func (m Metrics) FilePath() string {
	return "src/service/wiring"
}

func (m Metrics) FileName() string {
	return "metrics.go"
}

var metricsFile = `
package wiring

import (
	"gango/src/lib/monitor"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	StartTime = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: monitor.NameSpace,
		Name:      "gango_start_time",
		Help:      "gango start time",
	})
)
`
