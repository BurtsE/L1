package main

import "fmt"

func main() {
	var input = "ый"
	fmt.Println(swapString(input))

}

func swapString(s string) string {
	// Операция взятия по индексу возвращает некоторый байт строки. Т.к. символы unicode могут состоять из нескольких байт,
	// индексация по строке не будет корректной. Поэтому сначала конвертируем строку в слайс рун.
	var ans = []rune(s)
	for k := 0; k < len(ans)/2; k++ {

		// Обмен значений с протвоположных концов
		ans[k], ans[len(ans)-1-k] = ans[len(ans)-1-k], ans[k]
	}
	return string(ans)
}
