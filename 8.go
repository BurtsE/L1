package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	a, err := SetBit(50, 65)
	if err != nil {
		log.Fatal("err")
	}
	fmt.Println(a)
}

func SetBit(i int, val int64) (int64, error) {

	// Ошибка в случае, если номер бита превосходит максимально возможный
	if i > 63 {
		return val, errors.New("index is too big")
	}
	var k int64 = 1

	// Находим число, i-тый бит которого равен 1 при помощи битового свига
	for ; i > 1; i-- {
		k = k << 1
	}
	// Операция побитового или устанавливает в 1 те биты, которые отличны от нуля хотя бы в одном из чисел
	return val | k, nil
}
