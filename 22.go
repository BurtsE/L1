package main

import (
	"fmt"
	"math"
	"math/big"
)

// Для работы с большими числами используются функции из стандартного пакета math/big
func main() {
	var a = 1.
	var b = math.MaxFloat64 / 2.5

	fmt.Println(Add(a, b))
	fmt.Println(Subtract(a, b))
	fmt.Println(Multiply(a, b))
	fmt.Println(Divide(a, b))
}

func Add(a, b float64) *big.Float {
	x := big.NewFloat(a)
	y := big.NewFloat(b)
	z := big.NewFloat(0)
	z.Add(x, y)
	return z
}
func Subtract(a, b float64) *big.Float {
	x := big.NewFloat(a)
	y := big.NewFloat(b)
	z := big.NewFloat(0)
	z.Sub(x, y)
	return z
}
func Multiply(a, b float64) *big.Float {
	x := big.NewFloat(a)
	y := big.NewFloat(b)
	z := big.NewFloat(0)
	z.Mul(x, y)
	return z
}
func Divide(a, b float64) *big.Float {
	x := big.NewFloat(a)
	y := big.NewFloat(b)
	z := big.NewFloat(0)
	z.Quo(x, y)
	return z
}
