package app

import (
	"os"

	"github.com/Ioloman/migration-script/app/script/parallel"
	"github.com/Ioloman/migration-script/app/script/single"
	"github.com/urfave/cli/v2"
)

func Main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "batch-size",
				Aliases: []string{"b"},
				Value:   1000,
			},
			&cli.IntFlag{
				Name:    "print-every",
				Aliases: []string{"p"},
				Value:   10,
				Usage:   "how often to print the stats",
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "single",
				Usage: "sync execution",
				Action: func(ctx *cli.Context) error {
					Init()
					return single.Migrate(ctx.Int("batch-size"), ctx.Int("print-every"))
				},
			},
			{
				Name:  "parallel",
				Usage: "async execution",
				Action: func(ctx *cli.Context) error {
					Init()
					return parallel.Migrate(ctx.Int("batch-size"), ctx.Int("workers"), ctx.Int("print-every"))
				},
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:    "workers",
						Aliases: []string{"w"},
						Value:   5,
						Usage:   "how many concurrent workers to use",
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
