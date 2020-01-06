package commands

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rasterCmd.AddCommand(rasterPingCmd)
}

var rasterPingCmd = &cobra.Command{
	Use:   "ping",
	Short: "check the raster server status",
	Long:  "check the raster server status",
	Run: func(cmd *cobra.Command, args []string) {
		rasterReply := rasterClient.Ping()
		log.Debugf("Got ping reply: %v", rasterReply)
		println(rasterReply)
	},
}
