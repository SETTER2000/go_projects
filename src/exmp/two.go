package main

import (
	"fmt"
)

func joins(x, y string) (string, string) {
	return x, y
}

func main() {
	a, b := joins("Donna", "Rose") // Donna Rose
	fmt.Println(a, b)
}
