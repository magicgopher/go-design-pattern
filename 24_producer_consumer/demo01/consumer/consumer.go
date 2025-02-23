package consumer

import (
	"fmt"
	"github.com/magicgopher/go-design-pattern/24_producer_consumer/demo01/goods"
)

// Consumer 表示消费者
type Consumer struct {
	Name string
	ch   <-chan goods.Goods // 单向通道，只用于接收产品
}

// NewConsumer 创建一个新的消费者
func NewConsumer(name string, ch <-chan goods.Goods) *Consumer {
	return &Consumer{
		Name: name,
		ch:   ch,
	}
}

// Consume 开始消费产品
func (c *Consumer) Consume() {
	for product := range c.ch { // 从通道接收产品，直到关闭
		fmt.Printf("%s 正在消费: %+v\n", c.Name, product)
	}
	fmt.Printf("%s 消费完成\n", c.Name)
}
