package main

import (
	"fmt"
	"math/rand"
)

func main() {
	array := []int{5, 632, 4, 234, 14, 432, 23454, 4}
	fmt.Println(array)
	Qsort(array)
	fmt.Println(array)
}
func Qsort(A []int) {
	if len(A) <= 1 {
		return
	}
	// Деление массива на 2 части
	q := partition(A)
	// Для всех элеменов в результате процедуры partition справедливо: i < q   ==>   A[i] < A[q] ; i > q   ==>   A[i] < A[q]
	Qsort(A[:q])
	Qsort(A[q+1:])
}
func partition(A []int) int {
	r := len(A) - 1
	// Случайный выбор опорного элемента
	q := rand.Intn(len(A))
	A[q], A[r] = A[r], A[q]

	i := 0
	for j := range A {
		if A[j] < A[r] {
			A[i], A[j] = A[j], A[i]
			i++
		}
	}
	A[i], A[r] = A[r], A[i]
	return i
}
