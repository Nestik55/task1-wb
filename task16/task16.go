package main

import (
	"fmt"
	"math/rand"
)

func quicksort(a []int, l, r int) {
	// Остановка
	if l >= r {
		return
	}
	// Выбираем опорный элемент
	v := a[(l+r)/2]
	i := l
	j := r

	// Пробегаем по массиву, перемещаея меньшие опорного элемента влево, больше вправо
	for i <= j {
		for a[i] < v {
			i++
		}
		for a[j] > v {
			j--
		}
		if i >= j {
			break
		}
		// Меняем элементы местами
		a[i], a[j] = a[j], a[i]
		i++
		j--
	}

	// Вызываем рекурсивно на меньших отрезках
	quicksort(a, l, j)
	quicksort(a, j+1, r)
}

func main() {
	var a []int
	// Заполняем слайс ранодмными числами
	for i := 0; i < 20; i++ {
		a = append(a, rand.Intn(10000))
	}

	// Запускаем сортировку
	quicksort(a, 0, len(a)-1)

	fmt.Println(a)
}
