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
        Usage: "Custom tool for Laravel local development",
        Commands: []*cli.Command{
            {
				Name:  "init",
				Usage: "Initiate all required path and value",
				Action: func(cCtx *cli.Context) error {
					Nginx.Add()
					return nil
				},
			},
            {
				Name:  "nginx:add",
                Category: "NGINX",
				Usage: "Add a standard Laravel NGINX configuration files into sites-available directory",
				Action: func(cCtx *cli.Context) error {
					Nginx.Add()
					return nil
				},
			},
			{
				Name:  "nginx:remove",
                Category: "NGINX",
				Usage: "Remove a Laravel NGINX configuration files from sites-available directory",
				Action: func(cCtx *cli.Context) error {
					Nginx.Remove()
					return nil
				},
            },

            {
				Name:  "env:change",
                Category: "Laravel Environment",
				Usage: "Change .env values based on template",
				Action: func(cCtx *cli.Context) error {
					Env.Change()
					return nil
				},
			},
            {
				Name:  "env:add-template",
                Category: "Laravel Environment",
				Usage: "Add new a .env template",
				Action: func(cCtx *cli.Context) error {
					Env.Add_template()
					return nil
				},
			},
			{
				Name:  "env:remove-template",
                Category: "Laravel Environment",
				Usage: "Remove a .env template",
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