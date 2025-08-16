package main

import (
	"context"
	"time"
)

func Sleep(duration time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	<-ctx.Done()
	// fmt.Println("таймер все") - для проверки
}

func main() {
	Sleep(3 * time.Second)
}
