package queue

import (
	"sync"
	"testing"
)

func TestProducerConsumer(t *testing.T) {
	done := make(chan bool, 100)
	ch := make(chan int)
	go Producer(ch)
	go Consumer(ch, done)
	<-done
}

func TestProducerConsumer2(t *testing.T) {
	ch := make(chan int, 10)
	wg := &sync.WaitGroup{}
	//Producer2(ch,wg)
	//go Consumer2(ch, wg)
	go Consumer2(ch, wg)
	Producer2(ch, wg)
	wg.Wait()
}
