package components

import (
	"github.com/hexagram30/cli/components/config"
	rastercmpnt "github.com/hexagram30/raster/components"
	rastercfg "github.com/hexagram30/raster/components/config"
	logger "github.com/sirupsen/logrus"
)

// Base component collection
type Base struct {
	Config *config.Config
	Logger *logger.Logger
}

// BaseApp ...
type BaseApp struct {
	AppName     string
	AppAbbv     string
	ProjectPath string
	ConfigFile  string
}

// TestBase component that keeps stdout clean
type TestBase struct {
	Config *config.Config
}

// RasterClient ...
type RasterClient struct {
	rastercmpnt.BaseRPCClient
	RasterConfig *rastercfg.Config
}

// Default component collection
type Default struct {
	Base
	BaseApp
	RasterClient
}
