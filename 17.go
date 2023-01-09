package main

import (
	"fmt"
	"sort"
)

func main() {
	var input = []int{}

	fmt.Println(binarySearch(6, input))
	sort.Find()
}

// Бинарный поиск в отсортированном по возрастанию слайсе
func binarySearch(elem int, s []int) int {
	// Пустой слайс - выход
	if len(s) == 0 {
		return 0
	}
	// Границы слайса
	l := 0
	r := len(s) - 1
	for l < r {
		// Индекс центрального элемента
		m := int(uint(l+r) >> 1)
		// Смещение левой или правой границы
		if s[m] >= elem {
			r = m
		} else {
			l = m + 1
		}
	}
	// Если найденный элемент не равен искомому, значит искомый элемент отсутствует
	if s[l] != elem {
		return len(s)
	}
	return l
}
