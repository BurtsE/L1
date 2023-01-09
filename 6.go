package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	first()
	ch := make(chan int)
	go func() { ch <- 1 }()
	second(ch)
	third(true)
	time.Sleep(3 * time.Second)
}

// Горутина выполняет свою задачу и завершается
func first() {
	go func() {
		fmt.Println("First")
	}()
}

// Горутина останавливается при получениии сигнала из канала
func second(input chan int) {
	t := time.NewTimer(3 * time.Second)

	go func() {
		for {
			select {
			case _, ok := <-input:
				if !ok {
					return
				}
				fmt.Println("Second")
			case <-t.C:
				return
			}
		}
	}()
}

// Горутина завершается вызовом функции goexit()
func third(key bool) {
	go func() {
		if key {
			fmt.Println("Third")
			go func() {
				defer fmt.Println("ended")
				var a int
				for {
					a++
				}
			}()
			runtime.Goexit()
		}
	}()
}
