package queue

import (
	"fmt"
	"testing"
	"time"
)

func TestPublishSubscribe(t *testing.T) {
	p := NewPublisher(100*time.Microsecond, 10)
	defer p.Close()

	// 订阅者、订阅主题、主题发布消息
	tests := []struct {
		topic topic
		msg   interface{}
	}{
		0: {
			topic: "book",
			msg:   "Golang",
		},
		1: {
			topic: "book",
			msg:   "PHP",
		},
		2: {
			topic: "movie",
			msg:   "哈利波特",
		},
		3: {
			topic: "movie",
			msg:   "ring king",
		},
		4: {
			topic: "movie",
			msg:   "投名状",
		},
	}

	msgChans := make([]subscriber, len(tests))
	for k, tt := range tests {
		msgChans[k] = p.Subscribe(tt.topic)
		p.Publish(tt.topic, tt.msg)
	}
	for _, msgChan := range msgChans {
		p.Consume(msgChan, func(i interface{}) {
			fmt.Printf("receive message %v, subscriber %v \n", i, msgChan)
		})
	}
}
