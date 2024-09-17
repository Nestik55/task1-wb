package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := [5]int{2, 4, 6, 8, 10}
	sum := 0

	wg := new(sync.WaitGroup)
	mu := new(sync.Mutex)

	for i := 0; i < 5; i++ {
		// Увеличиваем кол-во работающих горутин на 1
		wg.Add(1)
		go func(wg *sync.WaitGroup, mu *sync.Mutex, num int) {
			// В последствии сообщаем, что задача выполнена
			defer wg.Done()
			// Блокируем доступ для других горутин
			mu.Lock()
			sum += num * num
			// Открываем доступ другим горутинам
			mu.Unlock()
		}(wg, mu, nums[i])
	}

	// Дожидаемся выполнения всех горутин
	wg.Wait()

	fmt.Println(sum)
}
