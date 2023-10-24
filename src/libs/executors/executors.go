package executors

import (
	"gango/utils"
	"path/filepath"
	"strings"
)

type Executors struct {
}

func (e Executors) WriteFolder(dir string) error {
	return utils.WriteFile(dir, filepath.Join(e.FilePath(), e.FileName()), strings.ReplaceAll(executorFile, "gango", dir))
}

func (e Executors) FilePath() string {
	return "/src/lib/executors"
}

func (e Executors) FileName() string {
	return "executors.go"
}

var executorFile = `
package executors

import (
	"errors"
	"os"
	"reflect"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)


var IdleWorkerError = errors.New("idle worker")

var (
	WorkersHistogram = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "worker_duration",
			Help:    "Per Type, in milli second",
			Buckets: []float64{1, 5, 10, 15, 20, 30, 50, 70, 100, 200, 500},
		},
		[]string{"worker_type"},
	)
	WorkerCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "worker_count",
		Help: "per status",
	},
		[]string{"worker_type", "status", "error_type"},
	)
)

func NewDefaultExecutor(worker IWorker) IExecutor {
	return &SimpleExecutor{
		Worker:      worker,
		WorkerSleep: 50 * time.Millisecond,
		IdleSleep:   5 * time.Second,
	}
}

func NewExecutor(worker IWorker, sleepTime time.Duration, idleTime time.Duration) IExecutor {
	return &SimpleExecutor{
		Worker:      worker,
		WorkerSleep: sleepTime,
		IdleSleep:   idleTime,
	}
}

type SimpleExecutor struct {
	Worker      IWorker
	WorkerSleep time.Duration
	IdleSleep   time.Duration
}

func (e *SimpleExecutor) Start(interrupt <-chan os.Signal) {
	isIdle := false
	workerName := reflect.ValueOf(e.Worker).Elem().Type().String()
	logger.Infof("Worker %s started", workerName)
	for {
		sleep := e.WorkerSleep
		if isIdle {
			sleep = e.IdleSleep
		}

		select {
		case <-interrupt:
			logger.Infof("Terminating %s.", workerName)
			return
		case <-time.After(sleep):
			isIdle = false
			timeStamp := time.Now()
			err := e.Worker.Loop()
			WorkersHistogram.WithLabelValues(reflect.TypeOf(e.Worker).String()).Observe(float64(time.Since(timeStamp).Milliseconds()))
			if err != nil {
				if err == IdleWorkerError {
					isIdle = true
					logger.Debugf("Sleeping %s.", workerName)
				} else {
					logger.Error(err)
					if errors.Unwrap(err) != nil {
						WorkerCounter.WithLabelValues(reflect.TypeOf(e.Worker).String(), "failed", errors.Unwrap(err).Error()).Inc()
					} else {
						WorkerCounter.WithLabelValues(reflect.TypeOf(e.Worker).String(), "failed",err.Error()).Inc()
					}
				}
			} else {
				WorkerCounter.WithLabelValues(reflect.TypeOf(e.Worker).String(), "succeed", "").Inc()
			}
		}
	}
}
`
