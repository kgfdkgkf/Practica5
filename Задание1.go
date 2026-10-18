package main

import(
	"fmt"
	"time"
)

func main() {
	result := make(chan string, 1)

	go func() {
		time.Sleep(3 * time.Second)
		result <- "Запрос выполнен"
	}()

	select {
	case m := <-result:
		fmt.Println(m)
	case <- time.After(2 * time.Second):
		fmt.Println("Время ожидания вышло")
	}
}