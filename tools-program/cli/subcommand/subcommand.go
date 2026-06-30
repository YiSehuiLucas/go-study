package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	// 创建命令集
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name: "add", // 命令名
				Aliases: []string{"a"}, // 命令别名
				Usage: "luzhengyang is the most handsome peoplen in the world",
				Action: func (cCtx *cli.Context) error{
					fmt.Println("added task:", cCtx.Args().First())
					return nil
				},
			},
			{
				Name: "complete",
				Aliases: []string{"c"},
				Usage: "complete a task on the list",
				Action: func (cCtx *cli.Context)error{
					fmt.Println("completed task:", cCtx.Args().First())
					return nil
				},
			},
			{
				Name: "template",
				Aliases: []string{"t"},
				Usage: "options for task templates",
				Subcommands: []*cli.Command{
					{
						Name: "add",
						Usage: "add a new template",
						Action: func(cCtx *cli.Context) error {
							fmt.Println("new task complate:", cCtx.Args().First())
							return nil
						},
					},
					{
						Name: "remove",
						Usage: "remove an existion template",
						Action: func(cCtx *cli.Context) error {
							fmt.Println("remove task complate:", cCtx.Args().First())
							return nil
						},
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}