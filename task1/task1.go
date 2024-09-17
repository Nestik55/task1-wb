package main

import "fmt"

// Создаю структуру Human
type Human struct {
	name string
	age  int
}

// Создаю структуру Action
type Action struct {
	Human // вложеная структура Human
	skill string
}

// Метод структуры Human
func (h *Human) getHuman() {
	fmt.Printf("My name is %s, I am %d age years old!\n", h.name, h.age)
}

// Метод структуры Action
func (act *Action) getAction() {
	fmt.Printf("Name: %s, Years: %d, Skill: %s\n", act.name, act.age, act.skill)
}

func main() {
	// Создаю экземпляр структуры Action
	act := Action{Human{"Valery", 17}, "programmer"}

	// Прямое обращение к методу Human
	act.getHuman()
	//My name is Valery, I am 17 age years old!

	act.getAction()
	//Name: Valery, Years: 17, Skill: programmer
}
