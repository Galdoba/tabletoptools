package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

const (
	programName = "mgt2"
)

func main() {
	app := cli.NewApp()
	app.Version = "v 0.0.0"
	app.Name = programName
	app.Usage = "Set of tools for MgT2"
	app.Description = "both interactive and automatic tools for Mongoose Traveller 2E tabletop game"
	app.Flags = []cli.Flag{}

	app.Before = func(c *cli.Context) error {
		return nil
	}
	//MAIN COMMANDS
	app.Commands = []*cli.Command{}

	app.After = func(c *cli.Context) error {
		return nil
	}
	args := os.Args
	if err := app.Run(args); err != nil {
		errOut := fmt.Sprintf("%v error: %v", programName, err.Error())
		println(errOut)
		os.Exit(1)
	}

}
