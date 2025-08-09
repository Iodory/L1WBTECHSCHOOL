package main

import "fmt"

type Human struct {
	Name    string
	Surname string
	Age     int
}

func (h *Human) Speak() {
	fmt.Printf("%s говорит\n", h.Name)
}
func (h *Human) Run() {
	fmt.Printf("%s бежит\n", h.Name)
}
func (h *Human) Sleep() {
	fmt.Printf("%s спит\n", h.Name)
}

type Action struct {
	Human
}

func main() {

}
