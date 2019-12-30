package common

import (
	"runtime"

	"github.com/geomyidia/util/filesystem"
)

// CallerPaths see the filesystem.CallerPaths function for details
func CallerPaths() *filesystem.CallerTree {
	return filesystem.CallerPaths(runtime.Caller(0))
}
