package middlwares

import (
	"gango/utils"
	"path/filepath"
)

var middleware = `
package middlewares

import (
	"context"
)

type AuthService struct {

}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (a *AuthService) Authenticate(ctx context.Context) (context.Context, error) {
	// TODO complete with your code
	return ctx, nil
}
`

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
