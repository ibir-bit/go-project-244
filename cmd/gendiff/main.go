package main

import (
	"fmt"
	"os"

	"code/cmd/gendiff/diff" // ← исправлено

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
				Usage:   "output format",
				Value:   "stylish",
			},
		},
		Action: func(c *cli.Context) error {
			if c.NArg() < 2 {
				return fmt.Errorf("требуется указать два файла для сравнения")
			}

			file1 := c.Args().Get(0)
			file2 := c.Args().Get(1)

			// Вызываем функцию из пакета diff
			_, _, err := diff.ParseFile(file1, file2)
			if err != nil {
				return err
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка: %v\n", err)
		os.Exit(1)
	}
}
