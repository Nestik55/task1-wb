package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Мапа из пакет sync обеспечивает безопастность без связки с мьютексом
	mp := &sync.Map{}

	// Запускаем горутины и используем метод Store
	for i := 0; i < 20; i++ {
		go mp.Store(i, i*100+1)
	}

	// Выводим значения мапы используя метод Load
	for i := 0; i < 20; i++ {
		v, ok := mp.Load(i)
		if !ok {
			fmt.Printf("Нет элемента с ключом: %d\n", i)
			continue
		}
		fmt.Println(v)
	}

	time.Sleep(5 * time.Second)
}
