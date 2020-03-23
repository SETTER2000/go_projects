package main

import (
	"bytes"
	"fmt"
	"strconv"
)

const (
	s  = "abcd"
	s1 = "абвг"
	x  = 100
)

//
//func getNextValidString() string {
//	return "Strings" + s
//}

func main() {
	var buffer bytes.Buffer
	fmt.Println(buffer)
	phrase := "vëtt og tþrt"
	fmt.Printf("string: \"%s\"\n", phrase)
	fmt.Println("index rune char bytes")
	//в каждой итерации программе становятся доступны индекс текущей позиции в строке
	//и кодовый пункт в этой позиции
	for index, char := range phrase {
		fmt.Printf("%-2d %U '%c' % X\n", index, char, char, []byte(string(char)))
	}
	/*for {
		if piece, ok := getNextValidString(); ok {
			buffer.WriteString(piece)
		} else {
			break
		}
	}*/

	bc := []rune(s)
	bc1 := []rune(s1)
	fmt.Printf("%s - преобразовать строку в срез кодовых пунктов Юникода: \n%v\n", s, bc)
	fmt.Printf("%s - преобразовать строку в срез кодовых пунктов Юникода: \n%v\n", s1, bc1)
	fmt.Printf("%s - преобразовать строку в срез байт без копирования: \n%v\n", s, []byte(s))
	fmt.Printf("%s - преобразовать строку в срез байт без копирования: \n%v\n", s1, []byte(s1))
	fmt.Printf("кол-во байт в строке %s: \n%v\n", s, len(s))
	fmt.Printf("кол-во байт в строке %s: \n%v\n", s1, len(s1))
	// Обратное преобразование где значение chars имеет тип []rune, или []int32,
	// а значение s будет иметь тип string
	fmt.Printf("Обратное преобразование среза %v в строку: \n%v\n", bc, string(bc))
	fmt.Printf("Обратное преобразование среза %v в строку: \n%v\n", bc1, string(bc1))
	t := fmt.Sprint(x)
	fmt.Printf("Cтроковое представление значения любого типа: \n %T %v => %T %v\n", x, x, t, t)
	r := strconv.Itoa(x)
	fmt.Printf("Cтроковое представление числа: \n %T %v \n", r, r)

}
