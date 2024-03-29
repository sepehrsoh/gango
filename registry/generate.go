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
	"gango/src/middlewares"
	"gango/src/services/configs"
	"gango/src/services/providers"
	"gango/src/services/wiring"
	"gango/utils"
)

type CustomConf struct {
	Providers []string
}

var (
	elastic  = "elastic"
	postgres = "postgres"
	redis    = "redis"
	grpc     = "grpc"
)

func NewCustomConfGenerator(config CustomConf) utils.Config {
	cnf := utils.NewDefaultConf()
	if utils.ArrayContain(redis, config.Providers) {
		cnf.Redis = true
	}
	if utils.ArrayContain(postgres, config.Providers) {
		cnf.Postgres = true
	}
	if utils.ArrayContain(elastic, config.Providers) {
		cnf.Elastic = true
	}
	if utils.ArrayContain(grpc, config.Providers) {
		cnf.Grpc = true
	}
	return cnf
}

func Generate(name string, conf utils.Config) {
	projectBase := base.BaseProject{}
	projectBase.Run(name)
	register := NewRegistry()
	registerLibrary(register)

	registerBase(register, conf)

	registerProviders(register, conf)

	registerWiring(register, conf)

	register.Register("src/hello-world", hello_world.NewHelloWord(utils.GetWireConfigs(conf)...))

	register.Run(name)

	finalize := base.FinalizeProject{}
	finalize.Run(name, conf)
}

func registerLibrary(register *Registry) {
	register.Register("lib/signals", misc.Signals{})
	register.Register("lib/monitor", monitor.Monitors{})
	register.Register("lib/logging", logging.Logging{})
	register.Register("lib/errs", errs.Errors{})
	register.Register("lib/executors/registry", executors.Registry{})
	register.Register("lib/executors/executors", executors.Executors{})
}

func registerBase(register *Registry, conf utils.Config) {
	register.Register("middlewares", middlewares.NewMiddleware(utils.GetWireConfigs(conf)...))
	register.Register("cmd/base", cmd.ProjectName{})
	register.Register("cmd/base/root", cmd.Root{})
	register.Register("cmd/base/serve", cmd.NewServe(utils.GetWireConfigs(conf)...))

	register.Register("service/configs/config", configs.NewConfig(utils.GetWireConfigs(conf)...))
}

func registerProviders(register *Registry, conf utils.Config) {
	if conf.Grpc {
		register.Register("service/providers/grpc", providers.Grpc{})
	} else {
		register.Register("service/providers/gin", providers.Gin{})
	}
	if conf.Redis {
		register.Register("service/providers/redis", providers.Redis{})
	}
	if conf.Elastic {
		register.Register("service/providers/elastic", providers.Elastic{})
	}
	if conf.Postgres {
		register.Register("service/providers/postgres", providers.Postgres{})
	}
}

func registerWiring(register *Registry, conf utils.Config) {
	register.Register("service/wiring/wiring", wiring.NewWiring(utils.GetWireConfigs(conf)...))
	register.Register("service/wiring/internal", wiring.NewInternal(utils.GetWireConfigs(conf)...))
	register.Register("service/wiring/metric", wiring.Metrics{})
	register.Register("service/wiring/service", wiring.NewService(utils.GetWireConfigs(conf)...))
}
