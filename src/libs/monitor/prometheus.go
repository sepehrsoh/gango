package monitor

import (
	"gango/utils"
	"path/filepath"
	"strings"
)

var monitors = `
package monitor

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"

	"gango/src/lib/logging"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

)

// TODO replace your name space
const NameSpace = "Gango"

var logger = logging.GetLogger("promethues")

type PrometheusMetricsServer struct {
	Url  string
	Port string
}

func (s PrometheusMetricsServer) Register(collector prometheus.Collector) {
	prometheus.MustRegister(collector)
}

func (s PrometheusMetricsServer) Start(interrupt <-chan os.Signal) {
	mux := http.NewServeMux()
	mux.Handle(s.Url, promhttp.Handler())

	addr := fmt.Sprintf(":%s", s.Port)
	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}
	go func() {
		logger.
			With("port", s.Port, "url", s.Url).
			Info("Create PrometheusMetricsServer")
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Fatalf("Monitor PrometheusMetricsServer Error: %v", err)
		}
	}()

	<-interrupt

	if err := server.Shutdown(context.Background()); err != nil {
		logger.Error(err)
	} else {
		logger.Info("Shout Down PrometheusMetricsServer")
	}
}
`

type Monitors struct{}

func (m Monitors) WriteFolder(dir string) error {
	return utils.WriteFile(dir, filepath.Join(m.FilePath(), m.FileName()), strings.ReplaceAll(monitors, "gango", dir))
}

func (m Monitors) FileName() string {
	return "prometheus.go"
}

func (m Monitors) FilePath() string {
	return "/src/lib/monitor"
}
