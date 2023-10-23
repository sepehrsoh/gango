package cmd

import (
	"gango/utils"
	"path/filepath"
	"strings"
)

type Root struct {
}

func (r Root) WriteFolder(dir string) error {
	return utils.WriteFile(dir, filepath.Join(r.FilePath(), r.FileName()), strings.ReplaceAll(rootFile, "gango", dir))
}

func (r Root) FilePath() string {
	return "/cmd/base"
}

func (r Root) FileName() string {
	return "root.go"
}

var rootFile = `
package base

import (
	"gango/src/lib/logging"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "gango",
		Short: "Create golang project",
	}
	logger =logging.GetLogger("cmd")
)

func Execute() {
	rootCmd.AddCommand(serveCmd)
	err := rootCmd.Execute()
	if err != nil {
		logger.Errorln(err)
	}
}

func init() {}
`
