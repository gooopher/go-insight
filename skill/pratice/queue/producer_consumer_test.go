package queue

import (
	"testing"
)

func TestProducerConsumer(t *testing.T) {
	done := make(chan bool, 100)
	ch := make(chan int)
	go Producer(ch)
	go Consumer(ch, done)
	<-done
}
