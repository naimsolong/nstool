package main

import (
    "log"
	"os"

	"naimsolong/nstool/nginx"
	"naimsolong/nstool/env"

    "github.com/urfave/cli/v2"
)

/*
 *
 *
 *
 */
func main() {
    app := &cli.App{
        Usage: "Custom tool for local development",
        Commands: []*cli.Command{
            {
				Name:  "nginx:add",
                Category: "NGINX",
				Usage: "add a nginx configuration files",
				Action: func(cCtx *cli.Context) error {
					Nginx.Add()
					return nil
				},
			},
			{
				Name:  "nginx:remove",
                Category: "NGINX",
				Usage: "remove a nginx configuration files",
				Action: func(cCtx *cli.Context) error {
					Nginx.Remove()
					return nil
				},
            },

            {
				Name:  "env:change",
                Category: "Laravel Environment",
				Usage: "change env values based on template",
				Action: func(cCtx *cli.Context) error {
					Env.Change()
					return nil
				},
			},
            {
				Name:  "env:add-template",
                Category: "Laravel Environment",
				Usage: "add new a env template",
				Action: func(cCtx *cli.Context) error {
					Env.Add_template()
					return nil
				},
			},
			{
				Name:  "env:remove-template",
                Category: "Laravel Environment",
				Usage: "remove a env template",
				Action: func(cCtx *cli.Context) error {
					Env.Remove_template()
					return nil
				},
            },
        },
    }

    if err := app.Run(os.Args); err != nil {
        log.Fatal(err)
    }
}