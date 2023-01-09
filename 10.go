package main

import (
	"fmt"
	"sort"
)

func main() {
	var input = []float64{1, 19}
	fmt.Println(Separate(input))
}

func Separate(s []float64) [][]float64 {
	if len(s) == 0 {
		panic("empty slice")
	}
	var ans = make([][]float64, 0)
	// Сортируем слайс, чтобы затем за один проход по нему разбить его на группы
	sort.Float64s(s)

	// Верхняя граница температуры для текущей группы
	limit := s[0] + 10

	// слайс значений текущей группы
	var step = make([]float64, 0)

	for i := 0; i < len(s); i++ {

		// Заполняем группу, пока не превысили границу
		if s[i] <= limit {
			step = append(step, s[i])
		} else {
			// При превышении добавляем группу к ответу, текущий элемент становится единственным в следующей группе. Сдвигаем границу относительно него
			limit = s[i] + 10
			ans = append(ans, step)
			step = make([]float64, 1)
			step[0] = s[i]
		}
	}
	// Последняя группа при выходе из цикла не добавлена в ответ
	ans = append(ans, step)
	return ans
}
