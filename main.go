package main

import (
	. "github.com/bravenewprince/Glcca/lexer"
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "Glcca"
	app.Usage = "Golang Compiler Compiler Again"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "file, f",
			Value: "data.txt",
			Usage: "the file you would like lexed",
		},
		cli.StringFlag{
			Name:  "grammar, g",
			Value: "grammar.glcca",
			Usage: "the file specifying your grammar",
		},
	}
	app.Action = func(c *cli.Context) {
		f := c.String("f")
		g := c.String("g")
		Lex(f, g)
	}
	app.Run(os.Args)
}
