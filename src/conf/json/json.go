package json

import (
	"encoding/json"
	"fmt"
	"os"
)

// Определение структурыЮ соответствующей значениям в JSON-файле
type configuration struct {
	Enabled bool
	Path    string
}

func main() {
	file, _ := os.Open("conf.json") // открытие конфигурационного файла
	defer file.Close()

	decoder := json.NewDecoder(file) // Извлечение JSON-значений в переменные
	conf := configuration{}
	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(conf.Path)
}
