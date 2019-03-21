package comsume

import (
	"fmt"
)

func consume(ch <-chan int, flag string) {
	for i := range ch {
		fmt.Println("Consumer: ", flag, " ---- ", i)
	}
}

func product(ch chan<- int) {
	for i := 0; i < 100000; i++ {
		fmt.Println("Product ---- ", i)
		ch <- i
	}
	ch <- -1
}

