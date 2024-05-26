package health

import (
	"fmt"

	"github.com/Galdoba/devtools/configmanager"
	"github.com/Galdoba/tabletoptools/definitions"
	"github.com/urfave/cli/v2"
)

func Check() *cli.Command {
	cmnd := &cli.Command{
		Name:        "health",
		Aliases:     []string{},
		Usage:       "Check program files",
		UsageText:   "taletoptools health",
		Description: "try to load config and all definition files",
		Action: func(c *cli.Context) error {
			_, err := configmanager.DefaultConfigPath(definitions.AppName)
			if err != nil {
				fmt.Println("config health: %v", err.Error())
				if err == configmanager.ErrNoConfig {
					fmt.Println("create default config?")
				}
			}

			return nil
		},
		Flags: []cli.Flag{},
	}
	return cmnd
}
