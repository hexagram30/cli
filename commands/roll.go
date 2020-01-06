package commands

import (
	"fmt"
	"os"
	"strconv"

	"github.com/crosscode-nl/partition"
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
		argCount := len(args)
		switch {
		case argCount == 0:
			log.Error("One or more arguments are required for rolling dice")
			println(cmd.Usage())
			os.Exit(1)
		case argCount >= 2 && argCount%2 != 0:
			log.Error("Various rolls require an even number of arguments")
			println(cmd.Usage())
			os.Exit(1)
		default:
			DispatchRolls(args)
		}
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
		formatRollOnce(result)
	case len(args) == 2:
		rollCount, err := strconv.ParseInt(args[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		results := diceClient.RollRepeated(die, rollCount)
		formatRollRepeated(results)
	default:
		var dice []string
		var counts []int64
		partition.ToFunc(len(args), 2, func(l int, h int) {
			roll := args[l:h]
			dice = append(dice, roll[0])
			count, err := strconv.ParseInt(roll[1], 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			counts = append(counts, count)
		})
		results := diceClient.RollVarious(dice, counts)
		formatRollVarious(results)
	}
}

func formatRollOnce(roll *api.DiceRoll) {
	fmt.Printf("%s:\n\t%d\n", roll.GetDiceType(), roll.GetResult())
}

func formatRollRepeated(rolls *api.DiceRepeatedRolls) {
	fmt.Printf("%s:\n\t%d\n", rolls.GetDiceType(), rolls.GetResults())
}

func formatRollVarious(rolls *api.DiceVariousRolls) {
	for _, roll := range rolls.GetResults() {
		fmt.Printf("%s:\n\t%v\n", roll.GetDiceType(), roll.GetResults())
	}
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

const rollLongDescription = `There are several formats supported for rolling dice:

    * a single roll
    * mulitple rolls of a single die
    * multiple rolls of various dice

  Supported dice include d4, d6, d8, d10, d12, d20, d100

  Examples:

    $ ./bin/hxgm30 roll d20
    $ ./bin/hxgm30 roll d20 5
    $ ./bin/hxgm30 roll d20 5 d12 2 d8 2 d6 6 d4 3
`
