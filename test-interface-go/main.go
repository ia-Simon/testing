package main

import "fmt"

type I interface {
	Foo()
	Bar()
}

type S struct {
	I
}

func (S) Foo() {
	fmt.Println("foo")
}

func main() {
	s := S{}

	fn := func(i I) {}

	fn(s)
	fn(s.I)

	s.I.Foo()
	s.Foo()
	s.I.Bar()
	s.Bar()
}
