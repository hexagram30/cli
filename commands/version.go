package commands

import (
	"encoding/json"
	"fmt"

	"github.com/geomyidia/util/version"
	"github.com/hexagram30/cli/common"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Version data as JSON",
	Long:  "Version data as JSON (for pretty formatting, pipe to `jq .`)",
	Run: func(cmd *cobra.Command, args []string) {
		version := versionToJSON(common.VersionData())
		fmt.Println(version)
	},
}

func versionToJSON(structData *version.ProjectVersion) string {
	jsonData, err := json.Marshal(structData)
	if err != nil {
		log.Error(err)
		log.Fatalf("Couldn't marshal: %v", structData)
	}
	return string(jsonData)
}
