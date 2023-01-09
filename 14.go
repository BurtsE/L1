package main

import (
	"fmt"
	"reflect"
)

func main() {
	var arg interface{}
	arg = make(chan<- string)
	arg = make(<-chan string)
	fmt.Println(defineType1(arg))
	fmt.Println(defineType1(3))

}
func defineType1(arg interface{}) string {
	switch reflect.ValueOf(arg).Kind() {
	case reflect.String:
		return "string"
	case reflect.Int:
		return "Int"
	case reflect.Bool:
		return "String"
	case reflect.Chan:
		return "Channel"
	default:
		return "unknown type"
	}

}
func defineType2(arg interface{}) (answer string) {
	// В случае, когда тип аргумента не относится к упомянутым 4, функция ChanDir из пакета reflect вызовет панику
	defer func() {
		if r := recover(); r != nil {
			answer = "unknown type"
		}
	}()
	// Интерфейс хранит в себе значение типа переменной, на которую указывает. Проверяем этот тип при помощи switch
	switch arg.(type) {
	case int:
		answer = "int"
	case string:
		answer = "string"
	case bool:
		answer = "bool"
	default:
		// Значения для каналов хранятся в связке с типом передаваемых в канал значений.
		// Поскольку канал может использоваться для передачи пользовательских типов, для идентификации изпользуется пакет reflect
		reflect.TypeOf(arg).ChanDir()
		answer = "channel"
	}
	return
}
