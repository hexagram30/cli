package commands

import (
	"github.com/hexagram30/cli/common"
	rasterclient "github.com/hexagram30/raster/client"
	rastercfg "github.com/hexagram30/raster/components/config"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(rasterCmd)
}

var rasterCmd = &cobra.Command{
	Use:   "raster",
	Short: "Operations against the Hexagram30 Raster server",
	Long:  longDescription,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		SetupRasterConnection()
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		TeardownRasterConnection()
	},
}

// SetupRasterConnection ...
func SetupRasterConnection() {
	rasterClient = rasterclient.New()
	rasterClient.Config = rastercfg.New()
	rasterClient.ProjectPath = common.CallerPaths().DotPath
	rasterClient.SetupConnection()
}

// TeardownRasterConnection ...
func TeardownRasterConnection() {
	rasterClient.Close()
}

const longDescription = `

This is a command line tool for querying the Hexagram30 raster service,
currently backed by RedixDB. This tool exercises a full path from client
to server and back again:
	1. string arguments (including JSON) in the terminal are converted to the
	   appropriate types;
	2. These are passed to the client library functions that are part of the
	   CLI tool;
	3. There, they are formatted appropriately for the gRPC service client
	   library and then passed as protobuf messages to the gRPC server;
	4. Upon receipt, the gRPC server executes the appropriate handlers, where
	   the data is converted to Go structs;
	5. For handlers that access the database, data is extracted from the
	   structs and used to form ScyllaDB queries;
	6. The database is queried, results parsed, and transformed into protobuf
	   payloads;
	7. These are received by the CLI client library and converted into string
	   data;
	8. Finally, the CLI tool writes these strings to stdout.

For commands that output JSON data, such as 'get ID' and 'version', you can get
nicely formatted reults by piping to jq, e.g.:

	$ ./bin/hxgm30 raster version | jq .

		{
		"version": "0.3.2",
		"buildDate": "2019-10-31T03:19:47Z",
		"gitCommit": "907d213",
		"gitBranch": "enhancement/78/true-subcommands",
		"gitSummary": "v0.3.2-14-g907d213-dirty"
		}
`
