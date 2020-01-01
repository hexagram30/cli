package commands

import (
	"encoding/json"

	diceapi "github.com/hexagram30/dice/src/golang/api"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	diceCmd.AddCommand(diceVersionCmd)
}

var diceVersionCmd = &cobra.Command{
	Use:   "version",
	Short: "the dice server version",
	Long:  "get the version of dice server that's currently being run",
	Run: func(cmd *cobra.Command, args []string) {
		versionReply := diceClient.Version()
		log.Debugf("Got version reply: %#v", versionReply)
		println(DiceRPCVersionToJSON(versionReply))
	},
}

// DiceRPCVersionToJSON ...
func DiceRPCVersionToJSON(structData *diceapi.VersionReply) string {
	jsonData, err := json.Marshal(structData)
	if err != nil {
		log.Error(err)
		log.Fatalf("Couldn't marshal: %v", structData)
	}
	return string(jsonData)
}
