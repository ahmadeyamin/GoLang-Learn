package main

import (
	"fmt"
)

type Speaker interface {
	Speak() bool
}

type Human struct {
	name string
}

type Cat struct {
	name string
}

type Dog struct {
	name string
}

func (s Human) Speak() bool {
	fmt.Println("Hello, I'm a human")
	return true
}

func (s *Cat) Speak() bool {
	fmt.Println("Hello, I'm a Cat")
	return true
}

func (s *Dog) Speak() bool {
	fmt.Println("Hello, I'm a Dog")
	return true
}

func main() {
	var t Speaker = Human{name: "John"}
	t.Speak()
}
