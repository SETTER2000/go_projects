package main

import (
	"fmt"
	"gopkg.in/urfave/cli.v1" // подключение пакета cli.go
	"os"
)

func main() {
	// Создание нового экземпляра приложения
	app := cli.NewApp()
	app.Name = "hello cli"
	app.Usage = "Print hello world"

	// Настройка флагов
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "name, n",
			Value: "World",
			Usage: "Who to say hello to",
		},
	}

	// Определение выполняемого действия
	app.Action = func(c *cli.Context) error {
		name := c.GlobalString("name")
		fmt.Printf("Hello %s!\n", name)
		return nil
	}
	app.Run(os.Args) // Запуск приложения
}
