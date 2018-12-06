// Программа выводит статистику по повторяющимся словам в файле
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func main() {
	var wg sync.WaitGroup // "Группа ожидания" для мониторина группы сопраграмм
	w := newWord()
	for _, f := range os.Args[1:] {
		wg.Add(1) // Сообщить "Группа ожидания" о ещё одной операции, по сути о ещё одной запущенной сопрограмме

		go func(file string) {
			if err := tallyWords(file, w); err != nil {
				fmt.Println(err.Error())
			}
			wg.Done() // Сообщить "Группе ожидания", что операция выполнена
		}(f)
	}
	/*
		"Группе ожидания" сообщает внешней программе (здесь main),
		чтоб она дождалась пока все сопрограммы закончатся, т.е. вызовут wg.Done()
	*/
	wg.Wait()


	fmt.Println("Words that appear more than once:")
	w.Lock() // Установка блокировки. Устранение эффета гонки (мьютексов).
	for word, count := range w.found {
		if count > 1 {
			fmt.Printf("%s: %d\n", word, count)
		}
	}
	w.Unlock() // Снятие блокировки
}

/*
Извлекаемые слова помещаются в структуру.
Можно также использовать тип данных map,
но применение структуры облегчит в дальнейшем реструктуризацию кода
*/
type words struct {
	sync.Mutex // Структура words теперь наследует мьютекс
	found map[string]int
}

func newWord() *words { // Создание нового экземляра слова
	return &words{found: map[string]int{}}
}

// Фиксируется кол-во вхождений этого слова. Если слово ещё на зафиксировано,
// добавить его. В противном случае увеличитм счётчик
func (w *words) add(word string, n int) {
	w.Lock() // Заблокировать объект, изменить и разблокировать
	count, ok := w.found[word]
	if !ok {
		w.found[word] = n
		return
	}
	w.found[word] = count + n
}

// Открытие файла, анализ содержимого и подсчёт найденных в нём слов.
// Функция копирования делает всё необхожимое.
func tallyWords(filename string, dict *words) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Сканер - полезный инструмент для подобного анализа файлов
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		dict.add(word, 1)
	}
	return scanner.Err()
}

/*func compress(filename string) {
	in, err := os.Open(filename) // Открыть исходный файл для чтения
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(filename + ".gz") // Открыть файл архива с расширением .gz и именем исходного файла
	if err != nil {
		return err
	}
	defer out.Close()

	//	Сжать данные и записать в соответствующий файл с помощью gzip.Writer
	//	Функция io.Copy выполняет необходимое копирование

	gzout := gzip.NewWriter(out)
	_, err = io.Copy(gzout, in)
	gzout.Close()
	return err
}*/
