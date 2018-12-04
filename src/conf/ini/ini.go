package main

import (
	"fmt"
	"gopkg.in/gcfg.v1" // Импорт пакета для работы с ini-файлами
)

func main() {
	/*
		Создание структуры для хранения конфигурационных значений
	*/
	config := struct {
		Section struct {
			Enabled bool
			Path    string
		}
		DB struct {
			Login    string
			Pass     string
			BaseName string
		}
	}{}
	err := gcfg.ReadFileInto(&config, "conf.ini")
	if err != nil {
		fmt.Println("Failed to perse config file: %s", err)
	}
	// Использование значений из ini-файла
	fmt.Println(config.Section.Enabled)
	fmt.Println(config.Section.Path)
	fmt.Println(config.DB.Login)
	fmt.Println(config.DB.Pass)
	fmt.Println(config.DB.BaseName)
}
