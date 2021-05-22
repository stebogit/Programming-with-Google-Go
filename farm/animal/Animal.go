package animal

import (
	"fmt"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type animal struct {
	locomotion string
	food       string
	noise      string
}

func New(locomotion, food, noise string) animal {
	return animal{locomotion, food, noise}
}

func (a animal) Eat() {
	fmt.Println(a.food)
}

func (a animal) Move() {
	fmt.Println(a.locomotion)
}

func (a animal) Speak() {
	fmt.Println(a.noise)
}
