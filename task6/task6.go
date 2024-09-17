package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func stopWtihChan(done <-chan struct{}) {
	for {
		select {
		// Если канал не пуст, завершаем выполнение
		case <-done:
			fmt.Println("End of stopWithChan!")
			return
			// В остальных случаях что-нибудь вычисляем
		default:
			time.Sleep(time.Second)
			fmt.Println("...StopWithChan is working...")
		}
	}
}

func stopWtihCloseChan(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		// Извлекаем число
		num, ok := <-ch
		// Если ok == false (в канале нет данных), завершаем выполнение горутины
		if !ok {
			fmt.Println("\nEnd of StopWithCloseChan")
			return
		}
		fmt.Print(num, " ")
	}
}

func stopWtihCtx(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println(time.Now())
	// Ожидание закрытия канали из контекста
	<-ctx.Done()

	fmt.Println(time.Now())
}

func main() {
	// 1 - Вариант
	// Создаем канал для сигнала о завершении
	done := make(chan struct{})

	//Запускаем функцию в горутине
	go stopWtihChan(done)

	// Ожидаем 5 секунд, записываем пустую структуру в канал для завершения выполения
	time.Sleep(5 * time.Second)
	done <- struct{}{}
	close(done)
	/*
				Output:
		...StopWithChan is working...
		...StopWithChan is working...
		...StopWithChan is working...
		...StopWithChan is working...
		...StopWithChan is working...
		End of stopWithChan!
	*/

	// 2 - вариант
	// Создаем канад и WaitGroup
	ch := make(chan int)
	wg := new(sync.WaitGroup)

	// Вызываем функцию в горутине
	go stopWtihCloseChan(ch, wg)

	wg.Add(1)
	for i := 0; i < 7; i++ {
		// Записываем данные в канал
		ch <- i
	}

	// Закрываем канал
	close(ch)
	// Ожидаем выполнение оставшихся горутин
	wg.Wait()

	/*
			Outout:
		0 1 2 3 4 5 6
		End of StopWithCloseChan
	*/

	//3 - Вариант
	// Создаем контекст с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	wg.Add(1)

	// Запускаем функцию в горутине
	go stopWtihCtx(ctx, wg)

	// Ожидаем выполнение оставшихся горутин
	wg.Wait()

	/*
						Output:
		2024-09-17 18:05:45.07427654 +0300 MSK m=+6.002256869
		2024-09-17 18:05:46.074423165 +0300 MSK m=+7.002402446
	*/

}
