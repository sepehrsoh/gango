package registry

import (
	"gango/src/base"
	"gango/src/base/cmd"
	"gango/src/libs/logging"
	"gango/src/libs/misc"
	"gango/src/libs/monitor"
	"gango/src/middlwares"
	"gango/src/services/configs"
	"gango/src/services/wiring"
)

func Generate(name string) {
	projectBase := base.BaseProject{}
	projectBase.Run(name)
	register := NewRegistry()

	register.Register("lib/signals", misc.Signals{})
	register.Register("lib/monitor", monitor.Monitors{})
	register.Register("lib/logging", logging.Logging{})
	register.Register("middlewares", middlwares.Middleware{})
	register.Register("cmd/base", cmd.ProjectName{})
	register.Register("cmd/base/root", cmd.Root{})
	register.Register("cmd/base/serve", cmd.Serve{})
	register.Register("service/configs/config", configs.Config{})
	register.Register("service/wiring/wiring", wiring.Wiring{})
	register.Register("service/wiring/internal", wiring.Internal{})

	register.Run(name)
}
