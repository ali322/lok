package main

import (
	"log"
	"os"

	"jav/lib"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Usage = "search files in mutiple external disk"
	app.Version = "0.0.1"

	app.Commands = []cli.Command{
		{
			Name:    "scan",
			Aliases: []string{"s"},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "path",
					Usage: "dest path",
				},
				cli.StringFlag{
					Name:  "remark",
					Usage: "disk remark",
				},
			},
			Usage:  "scan files in disk",
			Action: lib.Scan,
		},
		{
			Name:    "list",
			Aliases: []string{"ls"},
			Usage:   "list all mounted disk",
			Action:  lib.List,
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
