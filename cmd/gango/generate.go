package gango

import (
	"gango/registry"
	"github.com/spf13/cobra"
)

var (
	generateCmd = &cobra.Command{
		Run: func(cobraCmd *cobra.Command, args []string) {
			projectName, _ := cobraCmd.Flags().GetString("name")
			providers, _ := cobraCmd.Flags().GetStringArray("providers")
			if len(args) < 1 {
				logger.Panicln("./gango generate")
			}
			generate(projectName, registry.CustomConf{
				Providers: providers,
			})
		},
		Use:     "generate",
		Short:   "generate project",
		Long:    "generate new project with given providers (optional)",
		Example: "./gango generate [ -n  project-name ] [ -p [providers] ]  ",
	}
)

func generate(name string, conf registry.CustomConf) {
	registry.Generate(name, registry.NewCustomConfGenerator(conf))
}

func init() {
	generateCmd.Flags().StringP("name", "n", "", "project name")
	_ = generateCmd.MarkFlagRequired("name")
	generateCmd.Flags().StringArrayP("providers", "p", nil, "providers list")
}
