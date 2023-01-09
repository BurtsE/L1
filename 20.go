package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "asdf fdasf    ыfdaf"

	// Разбиваем строку на слова, создав слайс строк. Слова должны быть разделены пробелами
	words := strings.Split(input, " ")

	// Обращаем слайс
	for k := 0; k < len(words)/2; k++ {
		words[k], words[len(words)-1-k] = words[len(words)-1-k], words[k]
	}
	
	// Соединяем в единую строку
	fmt.Println(strings.Join(words, " "))
}
