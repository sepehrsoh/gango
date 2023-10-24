package middlwares

import (
	"gango/utils"
	"path/filepath"
)

type Middleware struct{}

func (m Middleware) FileName() string {
	return "middleware.go"
}

func (m Middleware) WriteFolder(dir string) error {
	return utils.WriteFile(dir, filepath.Join(m.FilePath(), m.FileName()), middleware)
}

func (m Middleware) FilePath() string {
	return "/src/middlewares"
}

var middleware = `
package middlewares

import (
	"github.com/gin-gonic/gin"
)

type Middleware struct {
}

func NewMiddleware() *Middleware {
	return &Middleware{}
}

func (m *Middleware) AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO complete with your logic
	}
}

`
