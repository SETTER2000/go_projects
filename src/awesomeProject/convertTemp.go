package main

import (
	"fmt"
	"github.com/SETTER2000/tempconv"
)

func main() {
	fmt.Println(tempconv.CToF(21))
	fmt.Println(tempconv.FToC(100))
	fmt.Println(tempconv.CToK(3))
	fmt.Println(tempconv.FToK(3))
	fmt.Println(tempconv.KToC(3))
}
