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

func main() {
	sobaka := &Dog{}
	coolSobaka := &AdaptDog{sobaka}

	fmt.Println(coolSobaka.Speak())
}
