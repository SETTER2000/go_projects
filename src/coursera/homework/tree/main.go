package main

import (
	"fmt"
	"io"
	"os"
	//"path/filepath"
	//"strings"
)

func dirTree(input io.Reader, output io.Writer, printFile bool) error {
	fmt.Println("i'm regular function")
}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}
