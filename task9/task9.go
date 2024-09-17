package main

import "fmt"

func main() {
	chInput := make(chan int)
	squares := make(chan int)

	// Первая горутина, которая записывает данные в канал
	go func() {
		for i := 0; i < 25; i++ {
			chInput <- i
		}
		close(chInput)
	}()

	// Горутина, которая вычисляет значения и записыает в другой канал
	go func() {
		for v := range chInput {
			squares <- v * 2
		}
		close(squares)
	}()

	// Вывод в главной go-подпрограмме
	for v := range squares {
		fmt.Print(v, " ")
	}
	fmt.Println()

}
