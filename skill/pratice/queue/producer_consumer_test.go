package queue

import (
	"fmt"
	"testing"
)

func TestProducerConsumer(t *testing.T) {
	done := make(chan bool, 100)
	ch := make(chan int)
	go Producer(ch)
	go Consumer(ch, done)
	<-done
}

func Producer(out chan int) {
	for i := 0; i < 1000; i++ {
		out <- i
	}
	close(out)
}

func Consumer(in chan int, done chan bool) {
	for v := range in {
		fmt.Println("receive message ", v)
	}
	done <- true
}
