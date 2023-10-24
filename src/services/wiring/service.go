package wiring

import (
	"gango/utils"
	"path/filepath"
	"strings"
)

type Service struct {
}

func (s Service) WriteFolder(dir string) error {
	return utils.WriteFile(dir, filepath.Join(s.FilePath(), s.FileName()), strings.ReplaceAll(serviceFile, "gango", dir))
}

func (s Service) FilePath() string {
	return "src/service/wiring"
}

func (s Service) FileName() string {
	return "service.go"
}

var serviceFile = `
package wiring

import (
	"fmt"

	"gango/src/service/providers"
	"gango/src/lib/logging"
	
	"github.com/go-redis/redis"
)

var logger = logging.GetLogger("service")

func (w *Wire) GetRedis() *redis.Client {
	if w.redis == nil {
		client, err := providers.NewRedisFromConfig(providers.RedisConfig{
			Server:   w.config.Redis.Server,
			Password: w.config.Redis.Password,
			DB:       w.config.Redis.DB,
		})
		if err != nil {
			panic(err)
		}
		w.redis = client
	}
	return w.redis
}

func (w *Wire) GetGinServer() func() {
	gin := providers.NewGinServer(w.GetAuthService())
	w.registerHandlers(gin)
	return func() {
		logger.With(
			"address", w.config.ServerAddress.Host,
			"port", w.config.ServerAddress.Port,
		).Infof("Start HTTP server")
		if err := gin.Run(
			fmt.Sprintf("%v:%v",  w.config.ServerAddress.Host, w.config.ServerAddress.Port)); err != nil {
			logger.Fatalf("listen: %s\n", err)
		}
	}
}

`
