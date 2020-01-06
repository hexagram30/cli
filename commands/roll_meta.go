package commands

import (
	"fmt"
	"os"
	"strconv"

	"github.com/crosscode-nl/partition"
	"github.com/hexagram30/dice/src/golang/api"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rollCmd.AddCommand(rollMetaCmd)
}

var rollMetaCmd = &cobra.Command{
	Use:   "meta",
	Short: "Make rolls with metadata (stats) using the dice server",
	Long:  rollMetaLongDescription,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		SetupDiceConnection()
	},
	Run: func(cmd *cobra.Command, args []string) {
		argCount := len(args)
		switch {
		case argCount == 0:
			log.Error("Two or more arguments are required for dice meta rolls")
			println(cmd.Usage())
			os.Exit(1)
		case argCount >= 2 && argCount%2 != 0:
			log.Error("Various meta rolls require an even number of arguments")
			println(cmd.Usage())
			os.Exit(1)
		default:
			DispatchMetaRolls(args)
		}
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		TeardownDiceConnection()
	},
}

// DispatchMetaRolls ...
func DispatchMetaRolls(args []string) {
	die := args[0]
	switch {
	case len(args) == 2:
		rollCount, err := strconv.ParseInt(args[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		results := diceClient.RollMetaRepeated(die, rollCount)
		formatMetaRollRepeated(results)
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
		results := diceClient.RollMetaVarious(dice, counts)
		formatMetaRollVarious(results)
	}
}

func formatMetaRollRepeated(metaRoll *api.MetaRoll) {
	roll := metaRoll.GetRoll()
	if roll != nil {
		fmt.Printf("%s:\n\t%d\n", roll.GetDiceType(), roll.GetResult())
	} else {
		rolls := metaRoll.GetRolls()
		fmt.Printf("%s:\n\t%v\n\t%v\n", rolls.GetDiceType(), rolls.GetResults(),
			metaRoll.GetStats())
	}
}

func formatMetaRollVarious(metaRolls *api.MetaRolls) {
	for _, metaRoll := range metaRolls.GetResults() {
		formatMetaRollRepeated(metaRoll)
	}
}

const rollMetaLongDescription = `Similar to the regular roll command, there is more than one supported format
  for rolling dice with roll metadata:

    * mulitple rolls of a single die
    * multiple rolls of various dice

  The metadata are statistics for the rolls, and include average, total roll
  count, highest roll, lowest roll, and a sum of all rolls.

  Note there is no metadata for a single roll (the statistics for that are not
  very interesting).

  Examples:

    $ ./bin/hxgm30 roll meta d20 5
    $ ./bin/hxgm30 roll meta d20 5 d12 2 d8 2 d6 6 d4 3
`
