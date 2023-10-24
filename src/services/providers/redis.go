package providers

import (
	"gango/utils"
	"path/filepath"
	"strings"
)

type Redis struct {
}

func (r Redis) WriteFolder(dir string) error {
	return utils.WriteFile(dir, filepath.Join(r.FilePath(), r.FileName()), strings.ReplaceAll(redisFile, "gango", dir))
}

func (r Redis) FilePath() string {
	return "src/service/providers"
}

func (r Redis) FileName() string {
	return "redis.go"
}

var redisFile = `
package providers

import (
	"gango/src/lib/logging"

	"github.com/go-redis/redis"
	"github.com/spf13/cast"
)

var logger = logging.GetLogger("providers")

func NewRedisFromConfig(conf RedisConfig) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     conf.Server,
        DB:       cast.ToInt(conf.DB),
		Password: conf.Password,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	logger.With("server" , conf.Server).Info("Connected to Redis.")
	return client, err
}

type RedisConfig struct {
	Server   string
	Password string
	DB       string
}

`
