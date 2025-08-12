package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Dog struct{}

func (d *Dog) Bark() string {
	return "Woof!"
}

type AdaptDog struct {
	Dog *Dog
}

func (a *AdaptDog) Speak() string {
	return a.Dog.Bark()
}

type Cat struct{}

func (c *Cat) Meow() string {
	return "Meow!!!!"
}

type AdaptCat struct {
	Cat *Cat
}

func (a *AdaptCat) Speak() string {
	return a.Cat.Meow()
}

func SomeFunc(s Speaker) {
	fmt.Println("Govorit:", s.Speak())
}

func main() {
	sobaka := &Dog{}
	coolSobaka := &AdaptDog{sobaka}

	kot := &Cat{}
	coolKot := &AdaptCat{kot}

	SomeFunc(coolSobaka)
	SomeFunc(coolKot)
}
