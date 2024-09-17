package main

import (
	"fmt"
	"unicode"
)

// Функция для проверки наличия неуникальных символов
func checkUniq(s string) bool {

	// Создадим мапу, что быстро проверять был ли уже такой символ
	mp := make(map[rune]bool)

	// Пробегаемся по строке
	for _, v := range s {
		// Приводим к нижнему регистру (для регистронезависимости)
		r := unicode.ToLower(v)

		// Если элемент уже был, возвращаем false
		if _, ok := mp[r]; ok {
			return false
		}

		// Запоминаем, что такой элемент уже был
		mp[r] = true
	}

	// Не нашли неуникальных элементов => все уникальные
	return true
}

func main() {
	arr := []string{"abcd", "abCdefAaf", "aabcd"}

	for _, v := range arr {
		fmt.Println(v, " - ", checkUniq(v))
	}

}
