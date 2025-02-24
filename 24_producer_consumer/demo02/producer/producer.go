package producer

import (
	"fmt"
	"github.com/magicgopher/go-design-pattern/24_producer_consumer/demo02/goods"
	"time"
)

// Producer 表示生产者
type Producer struct {
	Name string             // 名字
	ch   chan<- goods.Goods // 单向通道，只用于发送产品
}

// NewProducer 创建一个新的生产者
func NewProducer(name string, ch chan<- goods.Goods) *Producer {
	return &Producer{Name: name, ch: ch}
}

// Produce 开始生产产品
func (p *Producer) Produce(count int) {
	for i := 1; i <= count; i++ {
		product := goods.NewGoods(i, fmt.Sprintf("商品-%d", i))
		fmt.Printf("%s 正在生产: %+v\n", p.Name, product)
		p.ch <- product                    // 将产品发送到通道
		time.Sleep(500 * time.Millisecond) // 模拟生产耗时
	}
	close(p.ch) // 生产完成后关闭通道
}
