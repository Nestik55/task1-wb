package main

import "fmt"

// Функция дл получения отрезка, в пределах которого находится температура
func getSeg(t float32) int {
	return int(t/10) * 10
}

func main() {
	// Исходные данные
	temps := [...]float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	// Мапа для результата
	res := make(map[int][]float32)

	// Бежим по слайсу исходный данных
	for i := 0; i < len(temps); i++ {
		// Находим отрезок
		key := getSeg(temps[i])
		// Добавляем значение температуры для полученного отрезка
		res[key] = append(res[key], temps[i])
	}

	fmt.Println(res)
}
