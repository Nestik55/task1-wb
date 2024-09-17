package main

import (
	"fmt"
	"sync"
)

// Создаем структуру счетчика
type Counter struct {
	cnt int
	mu  sync.Mutex
}

func main() {
	// Создаем экземпляр счетчика
	counter := &Counter{
		cnt: 0,
		mu:  sync.Mutex{},
	}

	wg := new(sync.WaitGroup)

	// Запускаем 100000 горутин
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// Блокируем доступ к счетчику для других горутин
			counter.mu.Lock()
			// Увеличиваем счетчик горутин
			counter.cnt++
			// Открываем доступ другим горутинам
			counter.mu.Unlock()
		}()
	}

	// Дожидаемся выполнения
	wg.Wait()

	fmt.Println(counter.cnt)
}
