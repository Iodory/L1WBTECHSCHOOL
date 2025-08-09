package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func worker(id int, ch chan int) {
	for num := range ch {
		fmt.Printf("Воркер%d вывел: %d\n", id, num)
	}
}

func main() {
	var workers int

	fmt.Print("Введите количество воркеров: ")
	if _, err := fmt.Scan(&workers); err != nil {
		fmt.Println("Ошибка ввода", err)
		return
	}
	if workers <= 0 {
		fmt.Printf("Ошибка ввода воркеров, вы ввели: %d\n", workers)
		return
	}

	ch := make(chan int)

	for i := 1; i <= workers; i++ {
		go worker(i, ch)
	}

	stoper := make(chan os.Signal, 1)
	signal.Notify(stoper, os.Interrupt, syscall.SIGINT)

	go func() {
		var vvod int
		for {
			ch <- vvod
			vvod++
			// можно добавить time.Sleep(1 * time.Second) просто для удобного чтения, но не было написано что надо так что он тут)))
		}
	}()

	<-stoper
	fmt.Println("Закрытие канала")
	close(ch)
}
