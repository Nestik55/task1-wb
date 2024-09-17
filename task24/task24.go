package main

import (
	"fmt"
	"math"
)

// Структура точки
type Point struct {
	x float64
	y float64
}

// Конструктор для структуры точки
func newPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

// Геттер по X
func (p *Point) GetX() float64 {
	return p.x
}

// Геттер по Y
func (p *Point) GetY() float64 {
	return p.y
}

// Функция расчета расстояния
func dist(a, b *Point) float64 {
	x := a.GetX() - b.GetX()
	y := a.GetY() - b.GetY()
	return math.Sqrt(x*x + y*y)
}

func main() {
	a := newPoint(1.0, 2.0)
	b := newPoint(10.0, -1.0)

	fmt.Println(dist(a, b))
}
