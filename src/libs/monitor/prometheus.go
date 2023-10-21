package monitor

import (
	"gango/utils"
	"path/filepath"
)

var monitors = `
package monitor

import (
	"context"
	"errors"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

// TODO replace your name space
const NameSpace = "Gango"

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
		log.WithFields(log.Fields{
			"port": s.Port,
			"url":  s.Url,
		}).Info("Create PrometheusMetricsServer")
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Monitor PrometheusMetricsServer Error: %v", err)
		}
	}()

	<-interrupt

	if err := server.Shutdown(context.Background()); err != nil {
		log.Error(err)
	} else {
		log.Info("Shout Down PrometheusMetricsServer")
	}
}
`

type Monitors struct{}

func (m Monitors) WriteFolder(dir string) error {
	return utils.WriteFile(dir, filepath.Join(m.FilePath(), m.FileName()), monitors)
}

func (m Monitors) FileName() string {
	return "prometheus.go"
}

func (m Monitors) FilePath() string {
	return "/src/lib/monitor"
}
