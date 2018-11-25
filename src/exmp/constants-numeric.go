package main

import "fmt"

const (
	// Создайте огромное количество, переместив 1 бит влево на 100 мест.
	// Другими словами, двоичное число, равное 1, затем 100 нулей.
	Big = 1 << 100
	// Сдвиньте его снова на 99 мест, поэтому мы получим 1 << 1 или 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))   // 21
	fmt.Println(needFloat(Small)) // 0.2
	fmt.Println(needFloat(Big))   // 1.2676506002282295e+29
}
