package hello_world

import (
	"gango/utils"
	"path/filepath"
	"strings"
)

type HelloWord struct {
}

func (h HelloWord) WriteFolder(dir string) error {
	return utils.WriteFile(dir, filepath.Join(h.FilePath(), h.FileName()), strings.ReplaceAll(helloWorldFile, "gango", dir))
}

func (h HelloWord) FilePath() string {
	return "/src/hello-world"
}

func (h HelloWord) FileName() string {
	return "hello-world.go"
}

var helloWorldFile = `// Code generate by Gogang
package hello_world

import (
	"fmt"
	"net/http"

	"gango/src/lib/errs"

	"github.com/gin-gonic/gin"
)

type HelloWorldController struct {
}

func NewHelloWorldController() *HelloWorldController {
	return &HelloWorldController{}
}

func (d *HelloWorldController) helloWord(c *gin.Context) {
	_, err := c.Writer.Write([]byte(fmt.Sprint("hello world")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errs.ErrInternalServer})
		return
	}
}

func (d *HelloWorldController) RegisterRoutes(rg *gin.RouterGroup) {
	rg.GET("/hello-world", d.helloWord)
}

`
