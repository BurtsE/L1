package main

import (
	"fmt"
	"sync"
)

func main() {
	var input = []int{2, 4, 6, 8, 10}
	var sum int

	// Мьтекс для потокобезопасного выполнения расчетов
	mx := sync.Mutex{}

	// Создание группы для предотвращения печати до завершения расчетов
	wg := new(sync.WaitGroup)
	for _, i := range input {
		wg.Add(1)

		// Передаем горутине число, квадрат  которого нужно посчитать. Поскольку wg, mu в  бласти видимости горутины, их можно не передавать
		go func(ind int) {
			defer wg.Done()
			square := ind * ind
			mx.Lock()
			defer mx.Unlock()
			sum += square
		}(i)
	}
	wg.Wait()
	fmt.Println(sum)
}
