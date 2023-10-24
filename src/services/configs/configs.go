package configs

import (
	"gango/utils"
	"path/filepath"
	"strings"
)

type Config struct {
}

func (c Config) WriteFolder(dir string) error {
	return utils.WriteFile(dir, filepath.Join(c.FilePath(), c.FileName()), strings.ReplaceAll(configsFile, "gango", dir))
}

func (c Config) FilePath() string {
	return "src/service/configs"
}

func (c Config) FileName() string {
	return "configs.go"
}

var configsFile = `
package configs

import (
	"os"
)

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func NewConfigFromEnv() Config {
	return Config{
		ServerAddress: ServerAddress{
			Host: GetEnv("SERVER_HOST", "localhost"),
			Port: GetEnv("SERVER_PORT", "8000"),
		},
		Monitor: Monitor{
			Host: GetEnv("MONITOR_ROUT", "/metrics"),
			Port: GetEnv("MONITOR_PORT", "9090"),
		},
		Redis: Redis{
			Server:   GetEnv("REDIS_SERVER", "localhost"),
			DB:  	  GetEnv("REDIS_DB", "0"),
			Password: GetEnv("REDIS_PASSWORD", ""),
		},
	}
}

type Config struct {
	ServerAddress ServerAddress
	Monitor       Monitor
	Redis         Redis
}

type ServerAddress struct {
	Host string
	Port string
}

type Monitor struct {
	Host string
	Port string
}

type Redis struct {
	Server   string
	DB 	     string
	Password string
}
`
