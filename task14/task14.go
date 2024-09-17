package main

import (
	"fmt"
)

func main() {
	// Исходные данные
	arr := [4]any{9, "Hello", true, make(chan int)}

	// 1 способ
	// Пробегаем по arr
	for i := 0; i < 4; i++ {
		// Определяем тип
		switch arr[i].(type) {
		case int:
			fmt.Println(arr[i], "- type - int")
		case string:
			fmt.Println(arr[i], "- type - string")
		case bool:
			fmt.Println(arr[i], "- type - bool")
		default:
			fmt.Println("chan")
		}
	}

	// 2 способ
	// Пробегаем по arr
	for i := 0; i < 4; i++ {
		// Определяем тип
		switch arr[i].(type) {
		case int:
			fmt.Println(arr[i], "- type - int")
		case string:
			fmt.Println(arr[i], "- type - string")
		case bool:
			fmt.Println(arr[i], "- type - bool")
		default:
			fmt.Println("chan")
		}
	}

}
