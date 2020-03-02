package main

import "fmt"

func main() {
	v := 42 // change me!
	b := 0.123 + 0.214
	g := 0.867 + 0.5i
	fmt.Printf("value: %v type: %T\n", v, v) //value: 42 type: int
	fmt.Printf("value: %v type: %T\n", b, b) //value: 0.337 type: float64
	fmt.Printf("value: %v type: %T\n", g, g) //value: (0.867+0.5i) type: complex128
}
