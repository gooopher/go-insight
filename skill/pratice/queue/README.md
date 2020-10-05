# 队列

## 生产-消费模式
单元测试：
```shell script
go test -v skill/pratice/queue/producer_consumer.go skill/pratice/queue/producer_consumer_test.go
```


## 发布-订阅模式
发布 / 订阅（publish-subscribe）模型。生产-消费模式是将消息发送到一个队列，而发布-订阅模式是将消息发布给一个主题，发布者只管向主题发布消息，也不关心哪一个订阅者接收主题消息。订阅者和发布者可以在运行时动态添加，订阅者订阅了主题，就可以接收主题的消息，
实现消息的松耦合，摆脱了新业务接入需要重新起消费者多余的开发流程和系统复杂性增长。

常见场景：微博订阅、公众号订阅、短视频关注。

单元测试：
```shell script
go test -v skill/pratice/queue/publish_subscribe.go skill/pratice/queue/publish_subscribe_test.go
```
