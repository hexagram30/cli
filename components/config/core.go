package config

import (
	"fmt"
	"strings"

	logger "github.com/geomyidia/zylog/logger"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Configuration related constants
const (
	AppName         string = "cli"
	ConfigDir       string = "configs"
	ConfigFile      string = "app"
	ConfigType      string = "yaml"
	ConfigReadError string = "Fatal error config file"
)

func init() {
	viper.AddConfigPath(ConfigDir)
	viper.SetConfigName(ConfigFile)
	viper.SetConfigType(ConfigType)
	viper.SetEnvPrefix(AppName)

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.Set("Verbose", true)
	viper.AutomaticEnv()
	viper.AddConfigPath("/")

	err := viper.ReadInConfig()
	if err != nil {
		// log.Panic is not used here, since logging depends ...
		log.Panicf("%s: %s", ConfigReadError, err)
	}
}

// DiceRPCConfig ...
type DiceRPCConfig struct {
	Host string
	Port int
}

// DiceConfig ...
type DiceConfig struct {
	RPC *DiceRPCConfig
}

// Config ...
type Config struct {
	// We need to create a Dice config, since the Dice Go library is actually
	// part of a Clojure project with no Go configuration present.
	Dice *DiceConfig
	// The Raster client config is taken from the Raster project
	Logging *logger.ZyLogOptions
}

// New is a constructor that creates the full coniguration data structure
// for use by our application(s) and client(s) as an in-memory copy of the
// config data (saving from having to make repeated and somewhat expensive
// calls to the viper library).
//
// Note that Viper does provide both the AllSettings() and Unmarshall()
// functions, but these require that you have a struct defined that will be
// used to dump the Viper config data into. We've already got that set up, so
// there's no real benefit to switching.
//
// Furthermore, in our case, we're utilizing structs from other libraries to
// be used when setting those up (see how we initialize the logging component
// in ./components/logging.go, Setup).
func New() *Config {
	return &Config{
		Dice: &DiceConfig{
			RPC: &DiceRPCConfig{
				Host: viper.GetString("dice.grpc.host"),
				Port: viper.GetInt("dice.grpc.port"),
			},
		},
		Logging: &logger.ZyLogOptions{
			Colored:      viper.GetBool("logging.colored"),
			Level:        viper.GetString("logging.level"),
			Output:       viper.GetString("logging.output"),
			ReportCaller: viper.GetBool("logging.report-caller"),
		},
	}
}

// DiceConnectionString ...
func (c *Config) DiceConnectionString() string {
	return fmt.Sprintf("%s:%d", c.Dice.RPC.Host, c.Dice.RPC.Port)
}
