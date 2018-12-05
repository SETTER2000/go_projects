// Сжатие  произвольного кол-ва отдельных файлов. В Win 7 только по одному сжимает.  # gz exampledata/exm2.txt
package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

func main() {
	for _, file := range os.Args[1:] {
		er := compress(file)
		fmt.Println(er)
	}
}
func compress(filename string) error {
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

	/*
		Сжать данные и записать в соответствующий файл с помощью gzip.Writer
		Функция io.Copy выполняет необходимое копирование
	*/

	gzout := gzip.NewWriter(out)
	_, err = io.Copy(gzout, in)
	gzout.Close()
	return err
}
