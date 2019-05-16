package main

import (
	"log"
	"os"

	"lok/lib"

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
				cli.StringFlag{
					Name:  "ext",
					Usage: "stuff extnames",
				},
			},
			Usage:  "scan files in disk",
			Action: lib.Scan,
		},
		{
			Name:    "query",
			Aliases: []string{"q"},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "remark",
					Usage: "stuff remark",
				},
			},
			Usage:  "query stuffs by condition",
			Action: lib.Query,
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
