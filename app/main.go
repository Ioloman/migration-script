package app

import (
	"os"

	"github.com/Ioloman/migration-script/app/script/single"
	"github.com/urfave/cli/v2"
)

func Main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "single",
				Usage: "sync execution. specify batch size (--batch-size -b)",
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
				Action: func(ctx *cli.Context) error {
					return single.Migrate(ctx.Int("batch-size"), ctx.Int("print-every"))
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
