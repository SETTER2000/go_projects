package main

import (
	"fmt"
	"math"
	"strings"
)

const (
	Active BitFlag = 1 << iota
	Send
	Receive
)

type BitFlag int

func (flag BitFlag) String() string {
	var flags []string
	if flag&Active == Active {
		flags = append(flags, "Active")
	}
	if flag&Send == Send {
		flags = append(flags, "Send")
	}
	if flag&Receive == Receive {
		flags = append(flags, "Receive")
	}
	if len(flags) > 0 {
		return fmt.Sprintf("%d(%s)", int(flag), strings.Join(flags, "|"))
	}
	return "0()"
}
func main() {
	i := uint64(18446744073709551615)
	fmt.Println("Не самое большое положительное число в Go:")
	fmt.Printf("Type: %T %d \n", i, i)
	flag := Active | Send
	fmt.Println(BitFlag(0), Active, Send, Receive, flag)
	a, d := 2.0, 3.0
	r := math.Pow(d, a)
	fmt.Printf("Возведение в степень: %f", r)
	x := 10
	y := 4
	x &^= y
	fmt.Println(x)
	//bigDigits := [][]string{
	//	{"  000  ",
	//		" 0   0 ",
	//		"0     0",
	//		"0     0",
	//		"0     0",
	//		" 0   0 ",
	//		"  000  "},
	//	{" 1 ", "11 ", " 1 ", " 1 ", " 1 ", " 1 ", "111"},
	//}
	//fmt.Println(str[0] - '0')
	//fmt.Println(len(bigDigits))
}
