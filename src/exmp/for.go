/*
В Go есть только одна циклическая конструкция - цикл for.
Базовый цикл for имеет три компонента разделенные точкой с запятой:

блок инициализации: выполняется перед первой итерацией
условный блок: выполняется перед каждой итерацией
завершающий блок: выполняется в конце каждой итерации
Блок инициализации, как правило, содержит краткое объявление переменных;
переменные объявленные здесь доступны только в области видимости цикла for.
Цикл прекратит итерации, как только значение логического выражения в условном блоке будет false.
Замечание: В отличие от таких языков как C, Java или Javascript блоки конструкции for не
заключены в круглые скобки, а фигурные скобки { } всегда обязательны.*/
package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
		fmt.Println(sum)
	}
	fmt.Println(sum)
	// 0
	// 1
	// 3
	// 6
	// 10
	// 15
	// 21
	// 28
	// 36
	// 45
	// 45

}
