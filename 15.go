package main

import (
	"fmt"
)

var justString string

func someFunc() {
	v := createHugeString(1 << 10)

	// Строка может состоять из символов в кодировке Unicode, а также иметь длинну меньше 100
	// Исходный код может привести к потере данных или панике
	h := []rune(v)
	if len(h) >= 100 {
		justString = string(h[:100])
	} else {
		justString = string(h)
	}

}

func main() {
	someFunc()
	fmt.Println(justString)
}

func createHugeString(num int) string {
	str := make([]byte, num)
	for i := range str {
		str[i] = 90
	}
	return string(str)
}
