package gango

import (
	"gango/lib"
	"github.com/spf13/cobra"
)

var logger = lib.GetLogger("cmd")

var (
	rootCmd = &cobra.Command{
		Use:   "gango",
		Short: "Create golang project",
		Long:  `Create golang project.`,
	}
)

func Execute() {
	rootCmd.AddCommand(generateCmd)
	err := rootCmd.Execute()
	if err != nil {
		logger.Errorln(err)
	}
}

func init() {}
