package main

import (
	"fmt"
	"time"
)

func listenToChan(ch chan int) {
	for {
		i := <-ch
		fmt.Println("Got", i, "from channel")

		time.Sleep(1 * time.Second)
	}
}

func main() {
	ch := make(chan int, 10)

	go listenToChan(ch)

	for i := 0; i < 100; i++ {
		fmt.Println("Sending", i, "to channel")
		ch <- i
		fmt.Println("Sent", i, "to channel")
	}

	fmt.Println("All done. Closing channel.")

	close(ch)
}
