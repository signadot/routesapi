package router

import (
	"log/slog"
	"os"

	"github.com/signadot/routesapi/watched"
)

type Config struct {
	RouteServerAddr string
	Baseline        *watched.Baseline
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

func (c *Config) WithDebug(v bool) *Config {
	if v {
		c.levelVar.Set(slog.LevelDebug)
	} else {
		c.levelVar.Set(slog.LevelInfo)
	}
	return c
}
