package main

import (
	"fmt"
	"reflect"
	"strings"
)

func someAction(v []int8, b int8) {
	v[0] = 100
	v = append(v, b)
}

type m struct{}

func main() {
	b := strings.Builder{}
	b.WriteString()
	fmt.Println(reflect.TypeOf(m{}).Size())
	var a = []int8{1, 2, 3, 4, 5}
	someAction(a, 6)
	fmt.Println(a)
}
