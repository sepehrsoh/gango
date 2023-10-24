package cmd

import (
	"gango/utils"
	"path/filepath"
	"strings"
)

type Serve struct{}

func (s Serve) WriteFolder(dir string) error {
	return utils.WriteFile(dir, filepath.Join(s.FilePath(), s.FileName()), strings.ReplaceAll(serveFile, "gango", dir))
}

func (s Serve) FilePath() string {
	return "/cmd/base"
}

func (s Serve) FileName() string {
	return "serve.go"
}

var serveFile = `
package base

import (
	"gango/src/lib/misc"
	"gango/src/service/wiring"

	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serving gagngo service.",
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func serve() {
	wiringService := wiring.NewWire()
	wiringService.GetExecutorsRegistry().Start()
	httpServer := wiringService.GetGinServer()
	go httpServer()

	term := misc.CreateTerminateChannel()
	<-term
}
`
