package main

import "fmt"

func main() {
	sum := 0
	nums := [...]int{2, 4, 6, 8, 10}

	// создаём небуфиризированный канал
	chInput := make(chan int)

	// запускаем безымянную горутину
	go func() {
		defer close(chInput)
		for i := 0; i < len(nums); i++ {
			// читаем из канала
			num := <-chInput
			sum += num * num
		}
	}()

	// передаём числа в канал
	for i := 0; i < len(nums); i++ {
		chInput <- nums[i]
	}

	fmt.Println(sum)
}
