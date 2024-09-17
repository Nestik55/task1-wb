package main

import (
	"fmt"
)

func main() {

	nums := [5]int{2, 4, 6, 8, 10}
	// Выполняем задачу 5 раз в отдельных горутинах
	for i := 0; i < 5; i++ {
		<-calc(nums[i])
	}

}

func calc(num int) <-chan struct{} {
	// Создаем канал для синхронизации горутин
	done := make(chan struct{})
	// Выполняем задачу в горутине и закрываем канал
	go func() {
		fmt.Print(num*num, " ")
		close(done)
	}()

	return done
}
