package main

import (
	"fmt"
	"time"
)

func main() {
doIt:
	fmt.Println("1 first")
	go func() {
		fmt.Println("2 second")
	}()
	// Не забываем после {} ставить () - запуск анонимной функции
	go func() {
		fmt.Println("3 third")
	}()
	time.Sleep(150 * time.Millisecond)
	goto doIt
}
