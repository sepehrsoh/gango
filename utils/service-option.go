package utils

var DefaultOptions = ServiceOptions{
	WithRedis:    false,
	WithPostgres: false,
	WithGrpc:     false,
	WithElastic:  false,
}

type ServiceOptions struct {
	WithRedis    bool
	WithElastic  bool
	WithPostgres bool
	WithGrpc     bool
}

type Options interface {
	Apply(options *ServiceOptions)
}

type funcWireConfig struct {
	ops func(options *ServiceOptions)
}

func (w funcWireConfig) Apply(conf *ServiceOptions) {
	w.ops(conf)
}

func newFuncWireOption(f func(options *ServiceOptions)) *funcWireConfig {
	return &funcWireConfig{ops: f}
}

func WireWithRedis(enable bool) Options {
	return newFuncWireOption(func(options *ServiceOptions) {
		options.WithRedis = enable
	})
}

func WireWithElastic(enable bool) Options {
	return newFuncWireOption(func(options *ServiceOptions) {
		options.WithElastic = enable
	})
}

func WireWithPostgres(enable bool) Options {
	return newFuncWireOption(func(options *ServiceOptions) {
		options.WithPostgres = enable
	})
}

func WireWithGrpc(enable bool) Options {
	return newFuncWireOption(func(options *ServiceOptions) {
		options.WithGrpc = enable
	})
}

type Config struct {
	Redis    bool
	Elastic  bool
	Postgres bool
	Grpc     bool
}

func NewDefaultConf() Config {
	return Config{
		Redis:    false,
		Elastic:  false,
		Postgres: false,
		Grpc:     false,
	}
}

func GetWireConfigs(cnf Config) []Options {
	ops := make([]Options, 0)
	ops = append(ops, WireWithRedis(cnf.Redis))
	ops = append(ops, WireWithElastic(cnf.Elastic))
	ops = append(ops, WireWithPostgres(cnf.Postgres))
	ops = append(ops, WireWithGrpc(cnf.Grpc))
	return ops
}
