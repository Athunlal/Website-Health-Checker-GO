package app

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func Start() *cli.App {
	app := &cli.App{
		Name:  "Healthchecker",
		Usage: "A tiny tool that checks the given domain is down.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "domain",
				Aliases:  []string{"d"},
				Usage:    "Domain name to check.",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "port",
				Aliases:  []string{"p"},
				Usage:    "Port number to check.",
				Required: false,
			},
		},
		Action: func(c *cli.Context) error {
			port := c.String("port")
			if c.String("port") == "" {
				port = "80"
			}
			status := check(c.String("domain"), port)
			fmt.Println(status)
			return nil
		},
	}

	return app
}
