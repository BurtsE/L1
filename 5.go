package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	args := os.Args
	// Завершение, если не указано время работы программы
	if len(args) != 2 {
		log.Fatalf("Wrong input, Usage: main <processing time>\n")
	}
	N, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatal(err)
	}
	// Создание канала для записи значений. Закрывается создателем
	ch := make(chan interface{})
	defer close(ch)

	// Создание таймера
	t := time.NewTimer(time.Duration(N) * time.Second)

	// Запись в канал в отдельной горутине, сам канал в области её видимости.
	go func() {
		for {
			ch <- 1
			ch <- "sdf"
			ch <- 2 / 3
		}
	}()
	var i int
	for {
		select {
		// Пока таймеру не передан сигнал, выполняем чтение из канала
		case <-ch:
			fmt.Printf("signal: %d\n", i)
		// По истечении времени будет передан сигнал в канал таймера - завершение выполнения
		case <-t.C:
			return
		}
		i++
	}

}
