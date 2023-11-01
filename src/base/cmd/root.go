package cmd

import (
	"gango/utils"
)

type Root struct {
}

func (r Root) WriteFolder(dir string) error {
	return utils.EnrichTemplate(dir, r)
}

func (r Root) FilePath() string {
	return "/cmd/base"
}

func (r Root) FileName() string {
	return "root.go"
}

func (r Root) TemplateName() string {
	return "rootFile.tmpl"
}

func (r Root) TemplateData(name string) map[string]interface{} {
	return map[string]interface{}{
		"ProjectName": name,
	}
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
