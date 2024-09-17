package main

import (
	"fmt"
)

func changeBit(number *int64, bitNum int, bit bool) {

}

func main() {

	var number int64
	var bitNum int
	var bit bool

	fmt.Println("Введите последовательно число, номер бита и значение бита:")
	fmt.Scan(&number, &bitNum, &bit)
	if bit {
		// Ставим единицу в i-ный бит путём применения операции | (или)
		// между исходным числом и число состоящим из нулей и единицы в i-ом бите
		number = number | (1 << (bitNum - 1))
	} else {
		// Ставим единицу в i-ный бит путём применения операции,
		// затем интвертируем это число с помощью xor(^) и применяем операцию И(&)
		// таким образом везде будут 1, кроме i-го бита и после применения оперции к
		// number получим number c i-ным битом равным 1
		number = number&(1<<(bitNum-1)) ^ (1 << (bitNum - 1))
	}

	fmt.Println(number)
}
