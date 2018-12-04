package main

import (
	"fmt"
	"github.com/kylelemons/go-gypsy/yaml" // Импорт пакета YAML сторонних производителей
)

func main() {
	config, err := yaml.ReadFile("conf.yaml") // Чтение и анализ YAML-файла
	if err != nil {
		fmt.Println(err)
	}
	// Вывод значений из YAML-файла
	fmt.Println(config.Get("path"))
	fmt.Println(config.GetBool("enabled"))

}
