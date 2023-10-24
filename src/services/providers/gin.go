package providers

import (
	"gango/utils"
	"path/filepath"
	"strings"
)

type Gin struct {
}

func (g Gin) WriteFolder(dir string) error {
	return utils.WriteFile(dir, filepath.Join(g.FilePath(), g.FileName()), strings.ReplaceAll(ginFile, "gango", dir))
}

func (g Gin) FilePath() string {
	return "src/service/providers"
}

func (g Gin) FileName() string {
	return "gin.go"
}

var ginFile = `
package providers

import (
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
)

func NewGinServer(middleware gin.HandlerFunc) *gin.Engine {
	engine := gin.New()
	engine.Use(ginzap.Ginzap(logger.Desugar(), time.RFC3339, true))
	engine.Use(middleware)
	engine.Use(ginzap.RecoveryWithZap(logger.Desugar(), true))
	return engine
}

`
