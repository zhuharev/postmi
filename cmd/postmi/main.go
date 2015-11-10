package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "postmi"
	app.Usage = "fight the loneliness!"
	app.Action = run

	app.Run(os.Args)
}

func run(c *cli.Context) {

}
