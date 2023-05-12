package main

import (
	"log"
	"os"

	Init "naimsolong/nstool/init"
	Env "naimsolong/nstool/env"
	Nginx "naimsolong/nstool/nginx"

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
				Name:     "init",
				Usage:    "Initiate nstool configuration files",
				Action: func(cCtx *cli.Context) error {
					Init.Start(true)
					return nil
				},
			},
			{
				Name:     "nginx:show",
				Category: "NGINX",
				Usage:    "Show the existing NGINX configuration files",
				Action: func(cCtx *cli.Context) error {
					Nginx.List()
					return nil
				},
			},
			{
				Name:     "nginx:add",
				Category: "NGINX",
				Usage:    "Add a standard Laravel NGINX configuration files (require sudo)",
				Action: func(cCtx *cli.Context) error {
					Nginx.Add()
					return nil
				},
			},
			{
				Name:     "nginx:remove",
				Category: "NGINX",
				Usage:    "Remove a Laravel NGINX configuration files (require sudo)",
				Action: func(cCtx *cli.Context) error {
					Nginx.Remove()
					return nil
				},
			},

			{
				Name:     "env:copy",
				Category: "Laravel Environment",
				Usage:    "Copy .env from template or existing .env.example",
				Action: func(cCtx *cli.Context) error {
					Env.Copy()
					return nil
				},
			},
			{
				Name:     "env:value",
				Category: "Laravel Environment",
				Usage:    "Change .env values",
				Action: func(cCtx *cli.Context) error {
					Env.Change()
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
