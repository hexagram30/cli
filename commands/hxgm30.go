package commands

import (
	"github.com/hexagram30/cli/common"
	"github.com/hexagram30/cli/components"
	"github.com/hexagram30/cli/components/config"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// CLI ...
type CLI struct {
	components.Default
}

var cli = &CLI{}

// NewCLI ...
func NewCLI() *CLI {
	cli.AppName = config.AppName
	cli.AppAbbv = config.AppName
	cli.ProjectPath = common.CallerPaths().DotPath
	return cli
}

// // Close the gRPC connection
// func (cli *CLI) Close() {
// 	t.APIConn.Close()
// }

// Execute executes the root command, kicking off whatever setup is needed
// first.
func (c *CLI) Execute(args []string) error {
	c.Setup()
	return rootCmd.Execute()
}

// Setup ...
func (c *CLI) Setup() {
	cobra.OnInitialize(func() {
		// cli = c
	})
	cobra.AddTemplateFunc("Authors", func() string { return Authors })
	cobra.AddTemplateFunc("Copyright", func() string { return Copyright })
	cobra.AddTemplateFunc("Support", func() string { return Support })
	cobra.AddTemplateFunc("Website", func() string { return Website })
	rootCmd.SetHelpTemplate(helpTemplate)
}

// SetupConfiguration ...
func (c *CLI) SetupConfiguration() *config.Config {
	log.Debug("Updating configuration ...")
	log.Debugf("Using project path '%s' ...", c.ProjectPath)
	return config.New()
}
