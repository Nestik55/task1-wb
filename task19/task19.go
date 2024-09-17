package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	r := []rune(s)

	n := len(r)
	for i := 0; i < n/2; i++ {
		r[i], r[n-i-1] = r[n-i-1], r[i]
	}

	s = string(r)
	fmt.Println(s)
}
