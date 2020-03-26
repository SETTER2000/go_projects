package main

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
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
	g := "Саша"
	st := "Добро пожаловать в академию Go!"
	st2 := "Добро пожаловать в\u2028академию Go2!"
	var f, l = firstLastWord(st)
	var f2, l2 = firstLastParagraphWord(st2)
	fmt.Printf("Выбрать первый символ строки \"%v\": %v\n", g, string(g[0]))
	fmt.Printf("Выбрать первый символ строки \"%v\": %v\n", g, g[0:1])
	fmt.Printf("Выбрать первый кодовый пункт строки \"%v\": %v\n", g, g[0])
	fmt.Printf("Первое и последнее слово строки \"%v\": %v ... %v", st, f, l)
	fmt.Printf("Первое и последнее слово строки \"%v\": %v ... %v\n", st2, f2, l2)
	fmt.Printf("Логическое значение: %t %t\n", true, false)
	fmt.Printf("Логическое значение цифрами: %d %d\n", IntForBool(true), IntForBool(false))
	fmt.Printf("Восьмеричное представление числа 41: |%o|%#o|%# 8o|%#+ 8o|%+08o|\n", 41, 41, 41, 41, -41)
	i := 3931
	fmt.Printf("Шестнадцатеричное представления числа %v: |%x|%X|%8x|%08x|%#04X|0x%04X|\n", i, i, i, i, i, i, i)
	i = 596
	fmt.Printf("Десятичное представление числа %v: |$%d|$%06d|$%+06d|$%s|\n", i, i, i, i, Pad(i, 6, '*'))

	for _, x := range []float64{-.258, 7194.84, -60897162.0218, 1.500089e-8} {
		fmt.Printf("Форматирование вещественных значений: |%20.5e|%20.5f|%s|\n", x, x, Humanize(x, 20, 5, '*', ','))
	}
}
func Humanize(amount float64, width, decimals int, pad, separator rune) string {
	dollars, cents := math.Modf(amount)
	whole := fmt.Sprintf("%+.0f", dollars)[1:] // отбросить "+"
	fraction := ""
	if decimals > 0 {
		fraction = fmt.Sprintf("%+.*f", decimals, cents)[2:] // отбросить "+0"
	}
	sep := string(separator)
	for i := len(whole) - 3; i > 0; i -= 3 {
		whole = whole[:i] + sep + whole[i:]
	}
	if amount < 0.0 {
		whole = "-" + whole
	}
	number := whole + fraction
	gap := width - utf8.RuneCountInString(number)
	if gap > 0 {
		return strings.Repeat(string(pad), gap) + number
	}
	return number
}
func Pad(number, width int, pad rune) string {
	s := fmt.Sprint(number)
	gap := width - utf8.RuneCountInString(s)
	if gap > 0 {
		return strings.Repeat(string(pad), gap) + s
	}
	return s
}
func IntForBool(b bool) int {
	if b {
		return 1
	}
	return 0
}

func firstLastParagraphWord(s string) (string, string) {
	// Возвращает первое и последнее слово строки не зависимо от кол-ва пробелов в строке
	i := strings.IndexFunc(s, unicode.IsSpace)     // получаем индекс первой позиции пробела в строке
	firstWord := s[:i]                             // получить срез до первого пробела
	j := strings.LastIndexFunc(s, unicode.IsSpace) // получаем индекс последнего пробела в строке
	_, size := utf8.DecodeRuneInString(s[j:])      // возвращает кол-во байт в первом символе среза строки, начинающегося с последнего пробельного символа
	lastWord := s[j+size:]                         // получить срез последнего пробела
	return firstWord, lastWord
}
func firstLastWord(s string) (string, string) {
	// Возвращает первое и последнее слово строки
	i := strings.Index(s, " ")     //Получаем индекс первого пробела в строке
	firstWord := s[:i]             //Получить срез до первого пробела
	j := strings.LastIndex(s, " ") //Получаем индекс последнего пробела в строке
	lastWord := s[j+1:]            // Получить срез последнего пробела
	return firstWord, lastWord
}
