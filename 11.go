package main

import "fmt"

func main() {
	var set1 = map[int]bool{
		1: true,
		3: true,
		5: true,
		6: true,
	}
	var set2 = map[int]bool{
		1: true,
		3: true,
		6: true,
		7: true,
	}
	fmt.Println(intersect(set1, set2))
}

func intersect(a, b map[int]bool) map[int]bool {
	// map для пересечения множеств
	var ans = make(map[int]bool)

	// Все числа из 1 множества проходят проверку на наличие во 2, в случае совпадения значение добавляется в ответ
	for numb := range a {
		if _, ok := b[numb]; ok {
			ans[numb] = true
		}

	}
	return ans
}
