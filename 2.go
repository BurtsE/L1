package main

import (
	"fmt"
	"strings"
	"sync"
)

func main() {
	var input = []int{2, 4, 6, 8, 10}
	suqares(input)

}
func suqares(s []int) {
	// Создание группы для предотвращения печати до завершения расчетов
	wg := new(sync.WaitGroup)
	// Мьтекс для конкурентного (не параллельного) выполнения расчетов
	mu := sync.Mutex{}
	for i := range s {
		wg.Add(1)
		// Поскольку wg, mu в  области видимости горутины, их можно не передавать
		go func(num int) {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			s[num] *= s[num]
		}(i)
	}
	wg.Wait()
	fmt.Println(strings.Trim(fmt.Sprint(s), "[]"))

}
