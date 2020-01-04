package commands

import (
	"fmt"
	"os"
	"strconv"

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
		if len(args) == 0 {
			println(cmd.Usage())
			os.Exit(1)
		}
		DispatchRolls(args)
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		TeardownDiceConnection()
	},
}

// DispatchRolls ...
func DispatchRolls(args []string) {
	die := args[0]
	switch {
	case len(args) == 1:
		result := diceClient.RollOnce(die)
		println(formatRollOnce(result))
	case len(args) == 2:
		rollCount, err := strconv.ParseInt(args[1], 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		results := diceClient.RollRepeated(die, rollCount)
		println(formatRollRepeated(results))
	}
}

func formatRollOnce(result *api.DiceRoll) string {
	return fmt.Sprintf("%d", result.GetResult())
}

func formatRollRepeated(results *api.DiceRepeatedRolls) string {
	return fmt.Sprintf("%s: %d", results.GetDiceType(), results.GetResults())
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
