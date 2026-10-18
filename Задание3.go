package main

import "fmt"

func main() {
	ch := make(chan int)

	select {
	case n := <- ch:
		fmt.Println("Получено число: ", n)
	default:
		fmt.Println("Данных нет")
	}
}