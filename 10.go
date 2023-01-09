package main

import (
	"fmt"
	"sort"
)

func main() {
	var input = []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	fmt.Println(Separate2(input))
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

func Separate2(s []float64) map[int][]float64 {

	ans := make(map[int][]float64)
	for _, value := range s {

		// Деление в группы вида [0-10], [20-30]
		i := int(value) - (int(value) % 10)
		if _, ok := ans[i]; !ok {
			ans[i] = make([]float64, 0, 1)
		}
		ans[i] = append(ans[i], value)
	}

	return ans
}
