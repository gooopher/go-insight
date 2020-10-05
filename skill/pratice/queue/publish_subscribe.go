package queue

import (
	"fmt"
	"sync"
	"time"
)

type (
	subscriber chan interface{} // 订阅者
	topic      string           // 主题
)

func NewPublisher(timeout time.Duration, buffer int) *Publisher {
	return &Publisher{
		buffer:      buffer,
		timeout:     timeout,
		subscribers: make(map[subscriber]topic),
		wg:          make(map[subscriber]*sync.WaitGroup),
	}
}

type Publisher struct {
	m           sync.RWMutex                   // 读写锁
	buffer      int                            // 订阅队列缓冲区
	timeout     time.Duration                  // 发布超时时间
	subscribers map[subscriber]topic           // 订阅者
	wg          map[subscriber]*sync.WaitGroup // 等待组
}

// 订阅主题
func (p *Publisher) Subscribe(topic topic) chan interface{} {
	ch := make(chan interface{}, p.buffer)
	p.m.Lock()
	defer p.m.Unlock()
	p.subscribers[ch] = topic
	p.wg[ch] = &sync.WaitGroup{}
	return ch
}

// 取消订阅
func (p *Publisher) UnSubscribe(sub chan interface{}) {
	p.m.Lock()
	defer p.m.Unlock()
	delete(p.subscribers, sub)
	close(sub)
}

// 发布消息
func (p *Publisher) Publish(topic topic, v interface{}) {
	p.m.RLock()
	defer p.m.RUnlock()
	var wg sync.WaitGroup
	existTopic := false
	for sub, topic2 := range p.subscribers {
		if topic2 == topic {
			wg.Add(1)
			go p.sendToTopic(sub, topic, v, &wg)
			existTopic = true
		}
	}
	if !existTopic {
		err := fmt.Sprintf("%s", topic) + " topic not exist"
		panic(err)
	}
	wg.Wait()
}

// 发送主题，设置一定的超时
func (p *Publisher) sendToTopic(sub subscriber, topic topic, v interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	if topic == "" {
		return
	}
	select {
	case sub <- v:
	case <-time.After(p.timeout):
	}
	p.wg[sub].Add(1)
}

// 关闭主题chan
func (p *Publisher) Close() {
	p.m.Lock()
	defer p.m.Unlock()
	for sub, _ := range p.subscribers {
		delete(p.subscribers, sub)
		close(sub)
	}
}

func (p *Publisher) Consume(in chan interface{}, dealFunc func(interface{})) {
	go func() {
		for v := range in {
			dealFunc(v)
			p.wg[in].Done()
		}
	}()
	p.wg[in].Wait()
}
