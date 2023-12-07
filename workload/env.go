package workload

import "github.com/signadot/routesapi/watched"

func RouteserverAddr() string {
	return watched.GetRouteServerAddr()
}

func BaselineFromEnv() (*watched.Baseline, error) {
	return watched.BaselineFromEnv()
}
