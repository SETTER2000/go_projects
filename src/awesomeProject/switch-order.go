package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Когда суббота?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Сегодня.")
	case today + 1:
		fmt.Println("Завтра.")
	case today + 2:
		fmt.Println("Через два дня.")
	default:
		fmt.Println("Слишком далеко.")
	}
}
