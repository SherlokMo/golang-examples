package main

import "fmt"

type Animal interface {
	Speak() string
}

type cat struct{}
type dog struct{}
type JavaProgrammer struct{}

func (c cat) Speak() string {
	return "Meow"
}

func (d dog) Speak() string {
	return "Wooff"
}

func (j JavaProgrammer) Speak() string {
	return "Design Patterns!"
}

func talk(a Animal) {
	fmt.Println(a.Speak())
}

func main() {
	Animals := []Animal{cat{}, dog{}, JavaProgrammer{}}

	// using the generic talk function
	for _, animal := range Animals {
		talk(animal)
	}

	//using each struct method
	for _, animal := range Animals {
		fmt.Println(animal.Speak())
	}
}
