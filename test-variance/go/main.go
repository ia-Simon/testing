package main

type Animal interface {
	Breathe()
}

type Dog struct{}

func (Dog) Breathe() {}

func main() {
	var f func() Animal
	f = NewDog

	var g func(dog Dog)
	g = CheckAnimal
}

func NewDog() Dog {
	return Dog{}
}

func CheckAnimal(animal Animal) {}
