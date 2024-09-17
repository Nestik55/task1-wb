package main

import (
	"fmt"
	"sync"
)

func main() {
	mp := make(map[int]int)
	// Создали мапу, создаем WaitGroup и мьютекс
	wg := new(sync.WaitGroup)
	mu := new(sync.Mutex)

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(key int) {
			defer wg.Done()
			// Блокируем доступ для других горутин
			mu.Lock()
			// Записываем данные
			mp[key] = key*10 + 1
			// Даем доступ другим горутнам
			mu.Unlock()
		}(i)
	}

	// Ожидаем выполнения оставшихся горутин
	wg.Wait()

	// Выводим значения
	for i := 0; i < 20; i++ {
		fmt.Print(mp[i], " ")
	}
	fmt.Println()

}
