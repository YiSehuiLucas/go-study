package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"

)

func main() {
	var language string

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "language", // select language
				Value: "english",  // default language value
				Usage: "language for the greeting",
				Destination: &language, // 自动将 --language 的值写入 language 回填功能
			},
		},
		Action: func(cCtx *cli.Context) error {
			name := "default name(lucas)"
			if cCtx.NArg() > 0 {
				name = cCtx.Args().First()
			}
			if cCtx.String("language") == "spanish" {
				fmt.Println("Hola", name)
			} else {
				fmt.Println("hello", name)
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

	fmt.Println(language)
}
