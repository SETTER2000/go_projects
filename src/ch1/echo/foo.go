// Выводит аргументы задаваемые для программы
// os.Args[0:] - начиная прямо с команды
// os.Args[1:] - начиная с первого аргумента и заканчивая len(os.Args)
// т.е. до последнего аргумента os.Args[1:len(os.Args)], поэтому не пишется len(os.Args): os.Args[1:]
package main

import (
	"fmt"
	"os"
	//	"strings"
)

var x int

func main() {
	//	fmt.Println(strings.Join(os.Args[1:]," ")) // вывод только аргументов
	//	fmt.Println(os.Args[0:]) // просмотр для отладки без форматирования
	//  fmt.Println(strings.Join(os.Args[0:], " ")) // вывод начиная с команды и все последующие аргументы
	// 	fmt.Println(strings.Join(os.Args[1:],"\n")) // вывод начиная с команды и все последующие аргументы

	for x, arg := range os.Args[1:] {
		fmt.Println(x, arg)
	}

}
