package commands

import (
	diceclient "github.com/hexagram30/dice/src/golang/client"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(diceCmd)
}

var diceClient = &diceclient.Client{}

var diceCmd = &cobra.Command{
	Use:   "roll",
	Short: "Make rolls using the dice server",
	Long:  rollLongDescription,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		SetupDiceConnection()
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		TeardownDiceConnection()
	},
}

// SetupDiceConnection ...
func SetupDiceConnection() {
	log.Debug("Setting up connection to dice gRPC server ...")
	connStr := cli.Config.DiceConnectionString()
	diceClient = diceclient.New(connStr)
	diceClient.SetupConnection()
	log.Info(("Set up connection to dice gRPC server."))
}

// TeardownDiceConnection ...
func TeardownDiceConnection() {
	log.Debug("Closing connection to dice gRPC server ...")
	diceClient.Close()
	log.Info(("Closed connection to dice gRPC server."))
}

const rollLongDescription = `

TBD
`
