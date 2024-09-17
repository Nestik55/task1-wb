package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var a, b []int

	// Заполняем "Множества" рандомными числами до 15
	for i := 0; i < 15; i++ {
		a = append(a, rand.Intn(15))
		b = append(b, rand.Intn(15))
	}

	// Находим пересечение
	ans := intersect(a, b)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(ans)
}

func intersect(a []int, b []int) (res []int) {
	mpa := make(map[int]int)
	mpb := make(map[int]int)

	// Записываем значение 'a' в мапу для исключения повторяющихся значений
	// и более быстрого доступа в последствии
	for _, v := range a {
		mpa[v] = 1
	}

	// Аналогично  с 'b'
	for _, v := range b {
		mpb[v] = 1
	}

	// Ищем пересечение
	for k := range mpa {
		// Если элемент есть и в 'a' и в 'b', добавляем в ответ
		if _, ok := mpb[k]; ok {
			res = append(res, k)
		}
	}

	return res
}
