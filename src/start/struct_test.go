package main

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T){
	name := Cats{}
	fmt.Println(name.test())
	if name.test() == 5 {
		t.Error("Функция getName отдала неожиданное значение!" )
	}
}