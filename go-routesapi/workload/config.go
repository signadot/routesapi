package workload

import (
	"log/slog"
	"os"

	"github.com/signadot/routesapi/go-routesapi"
	"github.com/signadot/routesapi/go-routesapi/watched"
)

type Config struct {
	RouteServerAddr string
	Baseline        *watched.Baseline
	ClientInfo      *routesapi.ClientInfo
	Log             *slog.Logger
	levelVar        *slog.LevelVar
}

// EnvConfig attempts to read the config for a router
// from the environment.
func EnvConfig() (*Config, error) {
	baseline, err := BaselineFromEnv()
	if err != nil {
		return nil, err
	}
	addr := RouteserverAddr()
	levelVar := &slog.LevelVar{}
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: levelVar,
	}))
	return &Config{
		RouteServerAddr: addr,
		Baseline:        baseline,
		Log:             log,
		levelVar:        levelVar,
	}, nil
}

// WithDebug sets the log level of the logger created by [EnvConfig].  It has
// no effect on a Config not created by [EnvConfig].
func (c *Config) WithDebug(v bool) *Config {
	if c.levelVar == nil {
		return c
	}
	if v {
		c.levelVar.Set(slog.LevelDebug)
	} else {
		c.levelVar.Set(slog.LevelInfo)
	}
	return c
}
