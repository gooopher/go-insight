package queue

import (
	"fmt"
	"sync"
)

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

func Producer2(out chan int, wg *sync.WaitGroup) {
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		//go func(){
		out <- i
		//}()
	}
}

func Consumer2(in chan int, wg *sync.WaitGroup) {
	for v := range in {
		fmt.Println("receive message ", v)
		wg.Done()
	}
}
