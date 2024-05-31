package health

import (
	"github.com/Galdoba/tabletoptools/internal/mgt2/config"
	"github.com/urfave/cli/v2"
)

func Check() *cli.Command {
	cmnd := &cli.Command{
		Name:        "health",
		Aliases:     []string{},
		Usage:       "Check program files",
		UsageText:   "mgt2 files health",
		Description: "try to load config, assets and all definition files",
		Action: func(c *cli.Context) error {
			_, err := config.Load()
			if err != nil {
				return err
			}

			return nil
		},
		Flags: []cli.Flag{},
	}
	return cmnd
}
