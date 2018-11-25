package main

import (
	"flag"
	"fmt"
)

var name = flag.String("name", "World", "Имя для приветствия.")
var spanish, russian, deutsch, france, japan, china bool
var es, ru, fr, de, ja, ch = "испанский", "русский", "французский", "немецкий", "японский", "китайский"

func inStr(lan string) string {
	return "Использовать " + lan + " язык."
}

func init() {
	flag.BoolVar(&spanish, "spanish", false, inStr(es))
	flag.BoolVar(&spanish, "s", false, inStr(es))
	flag.BoolVar(&russian, "russian", false, inStr(ru))
	flag.BoolVar(&russian, "r", false, inStr(ru))
	flag.BoolVar(&france, "france", false, inStr(fr))
	flag.BoolVar(&france, "f", false, inStr(fr))
	flag.BoolVar(&deutsch, "deutsch", false, inStr(de))
	flag.BoolVar(&deutsch, "d", false, inStr(de))
	flag.BoolVar(&japan, "japan", false, inStr(ja))
	flag.BoolVar(&japan, "j", false, inStr(ja))
	flag.BoolVar(&china, "china", false, inStr(ch))
	flag.BoolVar(&china, "c", false, inStr(ch))
	flag.VisitAll(func(flag *flag.Flag) {
		format := "\t-%s: %s (Default: '%s')\n"
		fmt.Printf(format, flag.Name, flag.Usage, flag.DefValue)
	})
}

func main() {
	flag.Parse() // Парсинг флагов и присваивание значений в переменных
	switch {
	case spanish:
		fmt.Printf("Hola, %s!\n", *name) // Доступ к name как к указателю
	case russian:
		fmt.Printf("Здравствуйте, %s!\n", *name) // Доступ к name как к указателю
	case deutsch:
		fmt.Printf("Hallo, %s!\n", *name) // Доступ к name как к указателю
	case france:
		fmt.Printf("Bonjour, %s!\n", *name) // Доступ к name как к указателю
	case japan:
		fmt.Printf("こんにちは, %s!\n", *name) // Доступ к name как к указателю
	case china:
		fmt.Printf("你好, %s!\n", *name) // Доступ к name как к указателю
	default:
		fmt.Printf("Hello, %s!\n", *name)
	}
}
