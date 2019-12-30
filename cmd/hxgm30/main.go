package main

import (
	"os"

	"github.com/hexagram30/cli/commands"
	"github.com/hexagram30/cli/components/logging"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Create the client object and assign components to it
	cli := commands.NewCLI()
	// Bootstrap configuration and logging with defaults; this is to assist with
	// any debugging, e.g., logging output, etc.
	cli.Config = cli.SetupConfiguration()
	cli.Logger = logging.Load(cli.Config)
	err := cli.Execute(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
