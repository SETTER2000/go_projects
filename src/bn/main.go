package main

import "fmt"

/*
3
3
1
4
2
4
5
*/
func main() {
	three()
	three()
	defer five()
	defer four()
	one()
	defer two()
	four()
}

func one() {
	fmt.Println("1")
}
func two() {
	fmt.Println("2")
}
func three() {
	fmt.Println("3")
}
func four() {
	fmt.Println("4")
}
func five() {
	fmt.Println("5")
}
