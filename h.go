package main

import (
	"fmt"
)

// func someAction(v []int8, b int8) {
// 	v[0] = 100
// 	v = append(v, b)
// }

// type m struct{}

//	func main() {
//		b := strings.Builder{}
//		b.WriteString()
//		fmt.Println(reflect.TypeOf(m{}).Size())
//		var a = []int8{1, 2, 3, 4, 5}
//		someAction(a, 6)
//		fmt.Println(a)
//	}
func main() {
	slice := []string{"a", "a"}
	func(slice []string) {
		slice = append(slice, "a")
		slice[0] = "b"
		slice[1] = "b"
		fmt.Print(slice)
	}(slice)
	fmt.Print(slice)
}
