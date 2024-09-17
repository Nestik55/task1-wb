package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	var n int
	fmt.Print("Введите количество секунд: ")
	fmt.Scanln(&n)

	// Создадим контест с таймаутом, чтобы следить за временем
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(n)*time.Second)
	defer cancel()

	// Канал для чтения и записи
	ch := make(chan string)

	// Функция для чтения из канала
	go func() {
		// Как только канал закроется, функция завершится
		for s := range ch {
			fmt.Println(s, " ")
		}
	}()

	// Функция для записи в канал
	go func() {
		for {
			select {
			// Проверка на время (Если контест вернул канал, значит время вышло)
			case <-ctx.Done():
				// Закрываем канал и завершаем выполнение функции
				close(ch)
				return
			default:
				// Вводим данные и записываем их в канал
				var data string
				time.Sleep(time.Millisecond)
				fmt.Println("Введите данные: ")
				fmt.Scanln(&data)
				ch <- data
			}
		}
	}()

	// Произойдет, когда закончится время
	<-ctx.Done()
}
