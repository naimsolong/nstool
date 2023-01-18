package main

import (
    "log"
	"os"

	"naimsolong/nstool/nginx"
	"naimsolong/nstool/env"

	// "github.com/manifoldco/promptui"
    "github.com/urfave/cli/v2"
)

/*
 *
 *
 *
 */
func main() {
    app := &cli.App{
        Commands: []*cli.Command{
            {
				Name:  "nginx:add",
                Category: "NGINX",
				Usage: "add a nginx configuration files",
				Action: func(cCtx *cli.Context) error {
					Gonginx.Nginx_add()
					return nil
				},
			},
			{
				Name:  "nginx:remove",
                Category: "NGINX",
				Usage: "remove a nginx configuration files",
				Action: func(cCtx *cli.Context) error {
					Gonginx.Nginx_remove()
					return nil
				},
            },

            {
				Name:  "env:change",
                Category: "Laravel Environment",
				Usage: "change env values based on template",
				Action: func(cCtx *cli.Context) error {
					Goenv.Env_change()
					return nil
				},
			},
            {
				Name:  "env:add-template",
                Category: "Laravel Environment",
				Usage: "add new a env template",
				Action: func(cCtx *cli.Context) error {
					Goenv.Env_add_template()
					return nil
				},
			},
			{
				Name:  "env:remove-template",
                Category: "Laravel Environment",
				Usage: "remove a env template",
				Action: func(cCtx *cli.Context) error {
					Goenv.Env_remove_template()
					return nil
				},
            },
        },
    }

    if err := app.Run(os.Args); err != nil {
        log.Fatal(err)
    }
}