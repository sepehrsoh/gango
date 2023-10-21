package gango

import (
	"gango/registry"
	"gango/src/base"
	"gango/src/base/cmd"
	"gango/src/libs/misc"
	"gango/src/libs/monitor"
	"gango/src/middlwares"
	"github.com/spf13/cobra"
)

var (
	generateCmd = &cobra.Command{
		Run: func(cobraCmd *cobra.Command, args []string) {
			if len(args) < 1 {
				logger.Panicln("./quagmire migrate migrate-db users|all")
			}
			projectName := args[0]
			generate(projectName)
		},
		Use:     "generate",
		Short:   "generate project",
		Long:    "generate project",
		Example: "./quagmire migrate migrate-db users|all",
	}
)

func generate(name string) {

	projectBase := base.BaseProject{}
	projectBase.Run(name)
	register := registry.NewRegistry()

	register.Register("lib/signals", misc.Signals{})
	register.Register("lib/monitor", monitor.Monitors{})
	register.Register("middlewares", middlwares.Middleware{})
	register.Register("middlewares", cmd.ProjectName{})

	register.Run(name)
}
