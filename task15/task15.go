package main

/*
В данном коде есть две проблемы.

Во-первых, данная реализация приведет к утечке памяти, поскольку глобальная
переменная ссылается на срез, имеющий общую область памяти с бо́льшей строкой,
из-за чего сборщик мусора не сможет очистить неиспользуемую часть строки
после выхода из функции.

Во-вторых, в случае, если строка состоит из символов unicode,
операция взятия среза (v[:100]) вернет первые 100 байт, но не 100 символов.
Кроме того, после выполнения этой операции последний символ строки
может отображаться некорректно (если он занимал больше одного байта).
*/

/*

1. Если строка состоит из символов unicode,
операция взятия среза вернет 100 байт, а не 100 символов,
это, в свою очередь может привести к тому, что послдений символ
может отображаться некорректно, если он занимал больше одного байта
и так совпадёт по четности.

2. Сборщик мусора не очистит неиспользуемую часть строки, так как
глобальная переменная ссылается на срез, имеющий общую область памяти
с другой строкой
*/

var justString string

func createHugeString(len int) string {
	r := make([]rune, len)
	return string(r)
}

func someFunc() {
	// Будем использовать руны
	r1 := []rune(createHugeString(1 << 10))
	r2 := make([]rune, 100)

	// Скоптруем данные в новую область памяти
	copy(r2, r1[:100])

	justString = string(r2)
}

func main() {
	someFunc()
}
