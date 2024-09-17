package main

import (
	"fmt"
	"sync"
)

func main() {

	nums := [5]int{2, 4, 6, 8, 10}

	// Создаем WaitGroup
	wg := new(sync.WaitGroup)

	for _, v := range nums {
		// Добавляем к кол-ву работающий горутин 1
		wg.Add(1)
		go func(num int) {
			fmt.Print(num*num, " ")
			// Сообщаем WaitGroup что горутина завершила свою работу
			wg.Done()
		}(v)
	}

	// Дожидаемся выполенися всех горутин
	wg.Wait()
}
