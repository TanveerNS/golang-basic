package main

import (
	"fmt"
	"time"
)

func task(id int, ch chan<- string) {
	fmt.Printf("Task %d started\n", id)
	time.Sleep(2 * time.Second)
	ch <- fmt.Sprintf("Task %d completed", id)
}

func main() {
	ch := make(chan string)

	for i := 1; i <= 3; i++ {
		go task(i, ch)
	}

	for i := 1; i <= 3; i++ {
		fmt.Println(<-ch)
	}
}
