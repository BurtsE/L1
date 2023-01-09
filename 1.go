package main

import "fmt"

type Human []int64

func (h *Human) myName() {
	fmt.Println("i am human")
}
// Фактически наследования нет
// У структуры Action появляется атрибут с именем Human типа Human,
// имеющий метод myName. Вызвать этот метод можно, не указывая явно сам атрибут (a.Human.myName())
type Action struct {
	Human //Наследование методов от структуры Human
}

// Появление собственного метода оставит возможность обратиться к методу предка с помощью явного указания его имени

// func (a *Action) myName() {
// 	fmt.Println("I am action")
// }

func main() {
	var a Action
	a.myName()
	a.Human.myName()
}
