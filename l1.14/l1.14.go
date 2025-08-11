package main

import "fmt"

func main() {
	v := true
	detecter(v)
}

func detecter(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan int:
		fmt.Println("chan")
	default:
		fmt.Println("Неизвестный тип")
	}
}
