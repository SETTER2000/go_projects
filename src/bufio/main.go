// Выводит только уникальные строки
package main

import (
	"fmt"
)

const (
	_ = iota
	KB uint64 = 1 << (10*iota)
	MB
	d
)

func main() {
	//in := bufio.NewScanner(os.Stdin)
	fmt.Println( KB, MB, d)
	/*alreadySeen := make(map[string]bool)
	for in.Scan() {
		txt := in.Text()
		if _, found := alreadySeen[txt]; found {
			continue
		}
		alreadySeen[txt] = true
		fmt.Println(a, b, c, d)
	}*/
}
