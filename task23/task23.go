package main

import "fmt"

// Первый вариант, который меняет порядок, но работает за константу
func firstRemove(a []int, x int) []int {
	a[x] = a[len(a)-1]
	return a[:len(a)-1]
}

// Второй вариант, который сохраняет порядок, но работает за линию
func secondRemove(a []int, x int) []int {
	return append(a[:x], a[x+1:]...)
}

func main() {
	var a []int
	for i := 0; i < 10; i++ {
		a = append(a, i)
	}

	fmt.Println(a)
	a = firstRemove(a, 5)

	fmt.Println(a)
	a = secondRemove(a, 7)

	fmt.Println(a)
}
