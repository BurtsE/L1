package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(AllUnique("sdfg123вЫ"))
}

func AllUnique(input string) bool {

	// Уникальные символы сохраняем в map
	var seen = make(map[rune]bool)

	// Переводим всю строку в нижний регистр
	s := strings.ToLower(input)
	for _, sym := range s {
		// Запоминаем каждый символ
		if _, ok := seen[sym]; !ok {
			seen[sym] = true
		} else {
			// Символ уже был представлен - возвращаем false
			return false
		}
	}

	// Ни один символ не повторился - возвращаем true
	return true
}
