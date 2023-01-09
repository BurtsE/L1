package main

import (
	"fmt"
)

func main() {
	var a, b = 3, 4
	a, b = b, a // Обмен значений
	fmt.Println(a, b)
}
