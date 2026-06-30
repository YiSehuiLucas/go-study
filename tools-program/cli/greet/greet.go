package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "greet",
		Usage: "fight the loneliness",
		Action: func (cCtx *cli.Context) error {
			fmt.Println("Hello friend!")
			fmt.Printf("Hello %q", cCtx.Args().Get(1))
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}