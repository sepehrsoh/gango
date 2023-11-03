package gango

import (
	"gango/registry"

	"github.com/spf13/cobra"
)

var (
	generateCmd = &cobra.Command{
		Run: func(cobraCmd *cobra.Command, args []string) {
			if len(args) < 1 {
				logger.Panicln("./gango generate")
			}
			generate(args[0])
		},
		Use:     "generate",
		Short:   "generate project",
		Long:    "generate project",
		Example: "./gango generate [project-name]",
	}
)

func generate(name string) {
	registry.Generate(name)
}
