package commands

import (
	"fmt"

	"github.com/hexagram30/dice/src/golang/api"
	diceclient "github.com/hexagram30/dice/src/golang/client"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(rollCmd)
}

var diceClient = &diceclient.Client{}

var rollCmd = &cobra.Command{
	Use:   "roll",
	Short: "Make rolls using the dice server",
	Long:  rollLongDescription,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		SetupDiceConnection()
	},
	Run: func(cmd *cobra.Command, args []string) {
		DispatchRolls(args)
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		TeardownDiceConnection()
	},
}

// DispatchRolls ...
func DispatchRolls(args []string) {
	switch {
	case len(args) == 1:
		result := diceClient.RollOnce(args[0])
		println(formatRollOnce(result))
	}
}

func formatRollOnce(result *api.DiceRoll) string {
	return fmt.Sprintf("%d", result.GetResult())
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
