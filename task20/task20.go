package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reverse(s string) string {
	r := []rune(s)

	n := len(r)
	for i := 0; i < n/2; i++ {
		r[i], r[n-i-1] = r[n-i-1], r[i]
	}

	return string(r)
}

func main() {
	// Создаем буфер, чтобы ввести строку до символа перевода строки ('\n)
	in := bufio.NewReader(os.Stdin)
	// Вводим строку
	s, _ := in.ReadString('\n')

	// Создаем слайс с помощью сплита по пробелу
	arr := strings.Split(s[:len(s)-1], " ")

	// Разворачиваем каждую строку по отдельности
	for i := 0; i < len(arr); i++ {
		arr[i] = reverse(arr[i])
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

}
