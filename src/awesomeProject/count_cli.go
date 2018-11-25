package main

import (
	"fmt"
	"gopkg.in/urfave/cli.v1"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Usage = "Count up or down."
	app.Commands = []cli.Command{ // Определение одной или нескольких команд
		{
			Name: "up", ShortName: "u",
			Usage: "Count Up",
			Flags: []cli.Flag{ // Определение параметров команды
				cli.IntFlag{
					Name:  "stop, s",
					Usage: "Value to count up to",
					Value: 10,
				},
			},
			Action: func(c *cli.Context) error {
				start := c.Int("stop") // Получение параметра команды
				if start <= 0 {
					fmt.Println("Стоп не может быть отрицательным.")
				}
				for i := 1; i <= start; i++ {
					fmt.Println(i)
				}
				return nil
			},
		},
		{
			Name: "down", ShortName: "d",
			Usage: "Count Down",
			Flags: []cli.Flag{ // Определение параметров команды
				cli.IntFlag{
					Name:  "start, s",
					Usage: "Начать отсчет с",
					Value: 10,
				},
			},
			Action: func(c *cli.Context) error {
				start := c.Int("start")
				if start < 0 {
					fmt.Println("Start cannot be negative.")
				}
				for i := start; i >= 0; i-- {
					fmt.Println(i)
				}
				return nil
			},
		},
	}
	app.Run(os.Args)
}
