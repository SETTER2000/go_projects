package main

import (
	"fmt"
	"lengconv"
	"tempconv"
)

func main() {
	fmt.Println(lengconv.CmToFt(45))
	fmt.Println(tempconv.FToK(10))
	fmt.Println(tempconv.CToK(10))
	bob := Cats{"Mans", 6, .87}
	fmt.Println("Mans age is", bob.test())
}

/**
Структура данных.
Записываются название полей и тип данных
*/
type Cats struct {
	name      string
	age       int
	happiness float64
}

/**
Это метод структуры Cats (примерно как метод класса)
Метод называется test
Возвращает метод тип данных float64
Методом он становится благодаря безымянности  func, т.е. после func нет имени
следовательно это не функция. И второй момент, то что в параметре этот метод
имеет ссылку  на структуру (*Cats)
*/
func (cat *Cats) test() float64 {
	return float64(cat.age) * cat.happiness
}
