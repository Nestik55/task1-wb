package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	// Первый вариант
	// Классический вариант
	a, b = b, a
	fmt.Println(a, b)

	// Второй вариант
	// С помощью математики
	a += b
	b = a - b
	a -= b
	fmt.Println(a, b)
}
