package main

import (
	"fmt"
	"os"

	"github.com/Galdoba/tabletoptools/internal/mgt2/health"
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
	allCommands := []*cli.Command{}
	allCommands = append(allCommands, health.Commands()...)
	app.Commands = []*cli.Command{
		health.Check(),
	}

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

/*
mgt2 character new
mgt2 character print
mgt2 starship new
mgt2 freight stats
mgt2 hyperjump calculate
mgt2 hyperjump roll
mgt2 health

*/
