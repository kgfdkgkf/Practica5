package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	for i := 1; i <= 3; i++ {
		waitGroup.Add(1)

		go func(number int) {
			defer waitGroup.Done()
			fmt.Println("Горутина", number, "начала работу")

			time.Sleep(time.Duration(number) * time.Second)
			fmt.Println("Горутина", number, "завершила работу")
		}(i)
	}

	waitGroup.Wait()
	fmt.Println("Все горутины завершили работу")
}