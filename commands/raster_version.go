package commands

import (
	"encoding/json"
	"fmt"

	rasterapi "github.com/hexagram30/raster/api"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rasterCmd.AddCommand(rasterVersionCmd)
}

var rasterVersionCmd = &cobra.Command{
	Use:   "version",
	Short: "raster version data as JSON",
	Long:  "raster version data as JSON (for pretty formatting, pipe to `jq .`)",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(gRPCVersionToJSON(rasterClient.Version()))
	},
}

func gRPCVersionToJSON(structData *rasterapi.VersionReply) string {
	jsonData, err := json.Marshal(structData)
	if err != nil {
		log.Error(err)
		log.Fatalf("Couldn't marshal: %v", structData)
	}
	return string(jsonData)
}
