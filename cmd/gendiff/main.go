package main

import (
	"fmt"
	"log"
	"os"

	"code/pkg/gendiff"
	"code/pkg/parser"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "gendiff",
		Usage: "Compares two configuration files and shows a difference.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "format",
				Aliases: []string{"f"},
				Value:   "stylish",
				Usage:   "output format",
			},
		},
		Action: func(c *cli.Context) error {
			if c.NArg() != 2 {
				return cli.Exit("Error: two file paths are required", 1)
			}

			file1 := c.Args().Get(0)
			file2 := c.Args().Get(1)
			_ = c.String("format") // пока оставляем для будущего использования

			// Чтение и парсинг файлов
			data1, err := parser.ReadFile(file1)
			if err != nil {
				return cli.Exit(fmt.Sprintf("Failed to parse %s: %v", file1, err), 1)
			}

			data2, err := parser.ReadFile(file2)
			if err != nil {
				return cli.Exit(fmt.Sprintf("Failed to parse %s: %v", file2, err), 1)
			}

			// Генерация diff
			diff := gendiff.GenDiff(data1, data2)

			// Вывод результата
			fmt.Println(diff)

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
