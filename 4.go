package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
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

	// Канал для получения сигнала о завершении
	osCh := make(chan os.Signal, 1)
	defer close(osCh)
	signal.Notify(osCh, os.Interrupt, syscall.SIGTERM)

	// Канал для записи
	ch := make(chan interface{})
	defer close(ch)

	// Создание воркеров
	for i := 0; i < N; i++ {
		go func(i int) {

			// Пока канал открыт, будут производиться попытки чтения из него. Когда канал закроют, воркеры завершат выполнение
			for val := range ch {
				fmt.Printf("\nworker %d: %v\n", i, val)
				time.Sleep(3 * time.Second)
			}
		}(i)
	}

	// Запись данных в поток
	for {
		select {
		// При получени сигнала о завершении прекращаем работу
		case <-osCh:
			fmt.Println("End")

			//Go проходит по всем ожидающим на чтение или запись горутинам и разблокирует их
			//Все получатели получают значение по умолчанию для переменных заданного типа канала, а все отправители паникуют.
			return
		// when we do not have mes - we wait second and send info
		case <-time.After(time.Second):
			ch <- 1
			ch <- "sdf"
			ch <- 3.2
		}
	}

}
