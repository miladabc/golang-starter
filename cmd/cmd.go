package cmd

import (
	"os"

	"github.com/miladabc/golang-starter/internal/app"
	"github.com/rs/zerolog/log"
	cli "github.com/urfave/cli/v2"
)

func Execute() {
	cmd := &cli.App{
		Name:  "starter",
		Usage: "boilerplate for starting new golang projects",
		Commands: []*cli.Command{
			{
				Name:  "serve",
				Usage: "run http server",
				Action: func(*cli.Context) error {
					return app.StartHTTPServer()
				},
			},
		},
	}

	err := cmd.Run(os.Args)
	if err != nil {
		log.Fatal().Err(err).Msg("running cmd")
	}
}
