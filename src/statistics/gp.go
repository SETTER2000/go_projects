package main

import "fmt"

func main() {
	for i := 1e+07; i > 0; i-- {
		fmt.Println(i)
		if i == 4211817 {
			fmt.Printf("Соападение найдено: %v", i)
			return;
		}
	}
}
