package hal

import (
	"github.com/danryan/env"
)

// Config struct
type config struct {
	Name          string `env:"key=HAL_NAME default=hal"`
	Alias         string `env:"key=HAL_ALIAS"`
	AdapterName   string `env:"key=HAL_ADAPTER default=shell"`
	StoreName     string `env:"key=HAL_STORE default=memory"`
	Port          int    `env:"key=PORT default=9000"`
	LogLevel      string `env:"key=HAL_LOG_LEVEL default=info"`
	DebugEndpoint bool   `env:"key=HAL_DEBUG_ENDPOINT default=false"`
}

func newConfig() *config {
	c := &config{}
	env.MustProcess(c)
	return c
}
