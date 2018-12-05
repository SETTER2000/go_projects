// Эта программа использует сопрограмму для эхо-вывода в фоновом режиме, пока в основном режиме работает таймер.
package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	go echo(os.Stdin, os.Stdout) // Вызов функции echo как go-подпрограммы
	time.Sleep(30 * time.Second) // 30-секундная пауза
	fmt.Println("Time out.")
	os.Exit(0) // Выход из программы

}

func echo(in io.Reader, out io.Writer) { // Функция echo является обычной функцией
	io.Copy(out, in) // io.Copy скопирует данные из io.Reader в io.Writer
}
