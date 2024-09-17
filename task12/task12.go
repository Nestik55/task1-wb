package main

import (
	"fmt"
)

func main() {
	// Исходные данные
	sequence := [...]string{"cat", "cat", "dog", "cat", "tree"}

	// Будем использовать мапу, чтобы все значения сделать уникальными
	mp := make(map[string]struct{})

	// Пробегаем по исх. данным
	for _, v := range sequence {
		mp[v] = struct{}{}
	}

	// Результат
	for k := range mp {
		fmt.Println(k)
	}

}
