package main

import (
	"github.com/magicgopher/go-design-pattern/24_producer_consumer/demo02/consumer"
	"github.com/magicgopher/go-design-pattern/24_producer_consumer/demo02/goods"
	"github.com/magicgopher/go-design-pattern/24_producer_consumer/demo02/producer"
)

func main() {
	// 创建一个缓冲通道
	ch := make(chan goods.Goods, 5)

	// 初始化生产者和消费者
	p := producer.NewProducer("生产者-1", ch)
	c := consumer.NewConsumer("消费者-1", ch)

	// 启动生产者和消费者 goroutine
	go p.Produce(5) // 生产 5 个产品
	c.Consume()     // 主 goroutine 运行消费者
}
