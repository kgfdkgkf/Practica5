package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)

	go func() {
		for {
			select {
			case <- done:
				fmt.Println("Горутина остановлена")
				return
			case <- time.After(1 * time.Second):
				fmt.Println("Горутина работает")
			}
		}
	}()

	time.Sleep(5 * time.Second)
	close(done)
	fmt.Println("Программа завершена")
}