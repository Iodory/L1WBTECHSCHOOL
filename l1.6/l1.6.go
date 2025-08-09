package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

// боже как же я наговнокодил простите

func main() {
	num := 5
	stop := make(chan struct{})
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	ch := make(chan int, 1)
	ch <- 10
	CTRplusC := make(chan os.Signal, 1)
	signal.Notify(CTRplusC, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(CTRplusC)

	go yslovie(num)
	go emptyStruct(stop)
	go ctxStop(ctx)
	go Reader(ch)
	defer close(ch)
	go CTRLC(CTRplusC)
	go RunTime()

	defer close(stop)
	defer cancel()
	time.Sleep(10 * time.Second)
}

func emptyStruct(stop <-chan struct{}) {
	for {
		select {
		case <-stop:
			fmt.Println("stop")
			return
		default:
			fmt.Println("rabota")
			time.Sleep(5 * time.Second)
		}
	}
}

func yslovie(num int) {
	if num == 5 {
		fmt.Println("YSLOVIE STOP")
		return
	}
}

func ctxStop(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("CTX STOP")
			return
		default:
			fmt.Println("rabota")
			time.Sleep(5 * time.Second)
		}
	}
}

func RunTime() {
	defer fmt.Println("STOP RUNTIME")
	fmt.Println("rabota rabota....")
	runtime.Goexit()
}

func Reader(ch <-chan int) {
	for v := range ch {
		fmt.Printf("read %d\n", v)
		fmt.Println("read all, chan stop func end")
	}
}

func CTRLC(stop chan os.Signal) {
	<-stop
	fmt.Println("ctrl+c STOPPED")
}
