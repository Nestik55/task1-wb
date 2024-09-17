package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Функция читает данные из канала и выводит в Stdout
func work(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		x, ok := <-ch
		if !ok {
			return
		}
		fmt.Println(x)
	}
}

func main() {
	// n - кол-во воркеров
	var n int

	// определяем флаг, чтобы можно было задавать количество воркеров при запуске программы
	flag.IntVar(&n, "n", 3, "count of workers")
	flag.Parse()

	ch := make(chan int)

	// Канал для получения сигнала завершения
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT)

	// Создаем WaitGroup
	wg := sync.WaitGroup{}
	wg.Add(n)

	// Запускаем горутины для считывания данных
	for i := 0; i < n; i++ {
		go work(ch, &wg)
	}

	// Запускаем горутину записывающую рандомные числа
	go func() {
		for {
			select {
			case ch <- rand.Intn(100):
				time.Sleep(time.Second * 1 / 2)
			// Сигнал завершения
			case <-done:
				close(ch)
				return
			}
		}
	}()

	// Ожидаем завершения запущенных горутин
	wg.Wait()
}
