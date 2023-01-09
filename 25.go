package main

import (
	"fmt"
	"time"
)

func main() {
	Sleep3(4)

}

func Sleep1(sec int) {
	var sleepTime = sec * 2000000000
	for sleepTime > 0 {
		sleepTime--
	}
	fmt.Println("end")
}

func Sleep2(sec int) {
	t := time.NewTimer(time.Duration(sec) * time.Second)
	<-t.C
	fmt.Println("end")
}
func Sleep3(sec int) {

	//Пока не наступит определенное время, идет пустой цикл
	start := time.Now()

	//Конструктор сам распределит лишние секунды по другим полям
	end := time.Date(start.Year(), start.Month(), start.Day(), start.Hour(), start.Minute(), start.Second()+sec, start.Nanosecond(), start.Location())
	for time.Now().Before(end) {
	}
	fmt.Println("end")
}
