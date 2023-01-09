package main

import (
	"fmt"
	"main/channix"
	"math"
	"reflect"
)

func main() {
	// Инициализация точек
	p := channix.Point{}
	p.Point(1, 2)
	q := channix.Point{}
	q.Point(2, 3)

	d := Distance(p, q)
	fmt.Println(d)
}

func Distance(a, b channix.Point) float64 {

	// Предполагается, что названия полей, хранящих координаты, известно
	var x1, x2, y1, y2 float64

	// Определение значений координат
	_x1 := reflect.ValueOf(&a).Elem().FieldByName("x")
	_x2 := reflect.ValueOf(&b).Elem().FieldByName("x")
	_y1 := reflect.ValueOf(&a).Elem().FieldByName("y")
	_y2 := reflect.ValueOf(&b).Elem().FieldByName("y")

	// Определение типа переменных. Считаем одинаковым для всех полей
	t := _x1.Kind()

	if t <= 11 && t >= 2 { //Диапазон значений для целых чисел, в том числе и беззнаковых
		x1 = float64(_x1.Int())
		x2 = float64(_x2.Int())
		y1 = float64(_y1.Int())
		y2 = float64(_y2.Int())
	} else if t == 13 || t == 14 { //Диапазон значений для чисел с плавающей точкой
		x1 = _x1.Float()
		x2 = _x2.Float()
		y1 = _y1.Float()
		y2 = _y2.Float()
	} else {
		panic("not standart field type")
	}
	// Подсчет расстояния
	return math.Pow(math.Pow(x1-x2, 2)+math.Pow(y1-y2, 2), 0.5)
}
