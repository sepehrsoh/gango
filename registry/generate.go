package registry

import (
	"gango/src/base"
	"gango/src/base/cmd"
	hello_world "gango/src/hello-world"
	"gango/src/libs/errs"
	"gango/src/libs/executors"
	"gango/src/libs/logging"
	"gango/src/libs/misc"
	"gango/src/libs/monitor"
	"gango/src/middlwares"
	"gango/src/services/configs"
	"gango/src/services/providers"
	"gango/src/services/wiring"
)

func Generate(name string) {
	projectBase := base.BaseProject{}
	projectBase.Run(name)
	register := NewRegistry()

	register.Register("lib/signals", misc.Signals{})
	register.Register("lib/monitor", monitor.Monitors{})
	register.Register("lib/logging", logging.Logging{})
	register.Register("lib/errs", errs.Errors{})
	register.Register("lib/executors/registry", executors.Registry{})
	register.Register("lib/executors/executors", executors.Executors{})

	register.Register("middlewares", middlwares.Middleware{})

	register.Register("cmd/base", cmd.ProjectName{})
	register.Register("cmd/base/root", cmd.Root{})
	register.Register("cmd/base/serve", cmd.Serve{})

	register.Register("service/configs/config", configs.Config{})

	register.Register("service/providers/redis", providers.Redis{})
	register.Register("service/providers/gin", providers.Gin{})

	register.Register("service/wiring/wiring", wiring.Wiring{})
	register.Register("service/wiring/internal", wiring.Internal{})
	register.Register("service/wiring/metric", wiring.Metrics{})
	register.Register("service/wiring/service", wiring.Service{})

	register.Register("src/hello-world", hello_world.HelloWord{})

	register.Run(name)

	finalize := base.FinalizeProject{}
	finalize.Run(name)
}
