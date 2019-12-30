package commands

import (
	"strings"

	"github.com/hexagram30/cli/common"
	rasterclient "github.com/hexagram30/raster/client"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// Variables ...
var (
	Authors = strings.Join([]string{
		"Hexagram30 <hexagram30@cnbb.games>"}, "\n  ")
	Copyright = strings.Join([]string{
		"(c) 2019 Hexagram30"}, "\n  ")
	Support = "https://github.com/hexagram30/cli/issues"
	Website = "https://github.com/hexagram30/"
)

// var cli *CLI
var rasterClient *rasterclient.Client
var rootCmd = &cobra.Command{
	Use:   "hxgm30",
	Short: "The Hexagram30 CLI Tool",
	Long:  "",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		log.Tracef("Command: %#v", cmd)
		log.Debugf("Args: %#v", args)
		// cli.PostSetupPreRun()
	},
	Version: common.VersionString(),
}

var helpTemplate = `NAME
  {{.Name}} - {{.Short}}

USAGE{{if .Runnable}}
  {{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}
  {{.CommandPath}} [command]{{end}}

DESCRIPTION
  {{.Long}}{{if gt (len .Aliases) 0}}

ALIASES
  {{.NameAndAliases}}{{end}}{{if .HasExample}}

EXAMPLES
{{.Example}}{{end}}{{if .HasAvailableSubCommands}}

AVAILABLE COMMANDS{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}

FLAGS
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}

GLOBAL FLAGS
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}

TOPICS{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}

COPYRIGHT
  {{Copyright}}

AUTHORS
  {{Authors}}

WEBSITE
  {{Website}}

SUPPORT
  {{Support}}

Use "{{.CommandPath}} [command] --help" for more information about a command.{{end}}
`
