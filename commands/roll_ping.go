package commands

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rollCmd.AddCommand(rollPingCmd)
}

var rollPingCmd = &cobra.Command{
	Use:   "ping",
	Short: "check the dice server status",
	Long:  "check the dice server status",
	Run: func(cmd *cobra.Command, args []string) {
		pingReply := diceClient.Ping()
		log.Debugf("Got ping reply: %v", pingReply)
		println(pingReply)
	},
}
