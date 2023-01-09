package main

import "fmt"

func main() {
	// Исходный слайс
	a := []int8{1, 2, 3, 4, 5, 6}
	fmt.Println(RemoveI(0, a))
}

func RemoveI[T any](i int, input []T) []T {

	// Проверка на наличие индекса в слайсе
	if i < 0 || i >= len(input) {
		panic("list index out of range")
	}

	// Исключаем элемент. Используя срез из первой части слайса, заполняем его элементами из второй.
	return append(input[:i], input[i+1:]...)
}
