package animalsfight

import "fmt"

type Fighter struct {
	name  string
	power int
}

func NewFighter(name string, power int) *Fighter {
	return &Fighter{name, power}
}

func (f *Fighter) Introduce() {
	fmt.Println("Hi I am " + f.name)
}
