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
	// Ошибка, если число воркеров не указано при запуске
	if len(args) != 2 {
		log.Fatal("Wrong input. Usage:\n main <integer number of workers>\n")
	}
	N, err := strconv.Atoi(args[1])
	if err != nil {
		log.Printf("Usage:\n main <integer number of workers>\n")
		log.Fatal(err)
	}
	// Канал для записи
	ch := make(chan interface{})
	defer close(ch)
	// Создание воркеров
	for i := 0; i < N; i++ {
		go func(i int) {
			// Пока канал открыт, будут производиться попытки чтения из него
			for val := range ch {
				fmt.Printf("worker %d: %v\n", i, val)
			}
			time.Sleep(3 * time.Second)
		}(i)
	}

	// Запись данных в поток
	for i := 0; i < 5; i++ {
		ch <- 1
		ch <- "sdf"
		ch <- 3.2
		time.Sleep(1 * time.Second)
	}
	//Go проходит по всем ожидающим на чтение или запись горутинам и разблокирует их
	//Все получатели получают значение по умолчанию для переменных заданного типа канала, а все отправители паникуют.
	fmt.Println("end")
}
