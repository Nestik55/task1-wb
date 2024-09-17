package main

import (
	"fmt"
	"math/big"
)

func main() {

	// Cоздаём три переменных больших чисел
	var x, y, z = &big.Int{}, &big.Int{}, &big.Int{}

	// Инициализируем две переменный, а третью не трогаем
	x.Exp(big.NewInt(2), big.NewInt(23), nil)
	y.Exp(big.NewInt(2), big.NewInt(22), nil)

	// В третей переменной будет храниться результат вычислений первых двух
	z.Add(x, y)
	fmt.Printf("Cумма: %v\n", z.String())
	z.Mul(x, y)
	fmt.Printf("Умножение: %v\n", z.String())
	z.Sub(x, y)
	fmt.Printf("Вычитание: %v\n", z.String())
	z.Div(x, y)
	fmt.Printf("Деление: %v\n", z.String())
}
