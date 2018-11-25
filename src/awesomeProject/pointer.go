package main

import "fmt"

func point(p, k int) int {
	return p + k*96
}
func main() {
	a := 1
	k := new(int) // 0
	b := &a
	fmt.Println("Result:", point(*b, *k))
}
