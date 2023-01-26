package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	var str strings.Builder

	for i := 0; i < 100; i++ {
		str.WriteString("a")
	}
	a := str.String()
	fmt.Println(a)
	fmt.Println(reflect.TypeOf(a))
	b := fmt.Sprintf("%T", a)
	fmt.Println(b)
	// va, ex = string(str)
	// fmt.Println(va, ex)
}
