package watched

import "log/slog"

// Config provides the information to create a [Watched] or
// [BaselineWatched].
type Config struct {
	Addr string // address of the routeserver
	Log  *slog.Logger
}
