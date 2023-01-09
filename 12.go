package main

import "fmt"

func main() {
	var input = []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(makeSet(input))
}

func makeSet(s []string) []string {
	// Список слов при помощи map преобразуется во множество
	var set = make(map[string]bool)

	for _, word := range s {
		set[word] = true
	}

	// Множество возвращается обратно ввиде слайса
	var i int
	var ans = make([]string, len(set))
	for key := range set {
		ans[i] = key
		i++
	}
	return ans
}
