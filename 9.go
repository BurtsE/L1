package main

import (
	"fmt"
)

func main() {
	var a = []int{1, 2, 3, 4, 5}

	// Каналы закроются при завершении main
	// Канал для записи чисел
	input := make(chan int)

	//Канал для записи результатов
	output := make(chan int)

	// Массив и каналы находятся в области видимости горутин, их можно использовать без явной передачи.

	// Запись чисел в первый канал
	go func() {
		defer close(input) // Тот, кто пишет в канал, закрывает его
		for _, val := range a {
			input <- val
		}
	}()

	// Подсчет и запись результатов во 2 канал
	go func() {
		defer close(output) // Тот, кто пишет в канал, закрывает его
		for val := range input {
			output <- val * 2
		}
	}()

	// Чтение из 2 канала и вывод в stdout
	for val := range output {
		fmt.Println(val)
	}
}
