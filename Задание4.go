package main

import "fmt"

func main() {
	srochno := make(chan string, 2)
	obichno := make(chan string, 3)

	srochno <- "Срочная задача 1"
	srochno <- "Срочная задача 2"

	obichno <- "Обычная задача 1"
	obichno <- "Обычная задача 2"
	obichno <- "Обычная задача 3"

	finish := make(chan struct{})

	go func() {
		for i := 0; i < 5; {
			select {
			case task := <- srochno:
				fmt.Println("Обработана:", task)
				i++
				continue
			default:
			}

			select {
			case task := <- srochno:
				fmt.Println("Обработана:", task)
			case task := <- obichno:
				fmt.Println("Обработана:", task)
			}
			i++
		}
		close(finish)
	}()
	<-finish
	fmt.Println("Все задачи обработаны")
}