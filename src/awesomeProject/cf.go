// Cf конвертирует числовой аргумен в температуру
// по Цельсию, по Фаренгейту и по Кельвину
package main

import (
	"fmt"
	"github.com/SETTER2000/tempconv"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		k := tempconv.Kelvin(t)
		fmt.Printf("%s = %.3f°C\n%s = %.3f°K\n%s = %s\n%s = %s\n%s = %.3f°F\n%s = %.3f°C"+
			"\n", f, tempconv.FToC(f), f, tempconv.FToK(f),
			c, tempconv.CToF(c), c, tempconv.CToK(c), k, tempconv.KToF(k), k, tempconv.KToC(k))
	}
}
