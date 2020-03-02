package main

import "fmt"

func pud(sum int) (a float64, b int) {
	a = float64(sum) * 3 / 7
	b = sum + 10
	return
}

func main() {
	fmt.Println(pud(10)) // 4.285714285714286 20
}
