package main

import "time"

// Реализация sleep
func sleep(t time.Duration) {
	// В канал поступят данные после t
	// Он прочитается и функция завершится
	<-time.After(t)
}

func main() {
	// Пример использования собственного sleep
	sleep(7 * time.Second)
}
