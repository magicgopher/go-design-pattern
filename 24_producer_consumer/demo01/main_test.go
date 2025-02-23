package main

import (
	"github.com/magicgopher/go-design-pattern/24_producer_consumer/demo01/goods"
	"github.com/magicgopher/go-design-pattern/24_producer_consumer/demo01/producer"
	"testing"
	"time"
)

// TestProducerProduce 测试生产者的生产功能
func TestProducerProduce(t *testing.T) {
	ch := make(chan goods.Goods, 5)
	p := producer.NewProducer("测试生产者", ch)

	// 启动生产者
	go p.Produce(3)

	// 验证生产的产品
	expected := []goods.Goods{
		goods.NewGoods(1, "商品-1"),
		goods.NewGoods(2, "商品-2"),
		goods.NewGoods(3, "商品-3"),
	}

	i := 0
	for product := range ch {
		if product.ID != expected[i].ID || product.Value != expected[i].Value {
			t.Errorf("期望 %+v, 得到 %+v", expected[i], product)
		}
		i++
	}

	if i != 3 {
		t.Errorf("期望生产 3 个产品, 实际得到 %d 个", i)
	}
}

// TestProducerChannelClose 测试生产者关闭通道
func TestProducerChannelClose(t *testing.T) {
	ch := make(chan goods.Goods, 5)
	p := producer.NewProducer("测试生产者", ch)

	go p.Produce(2)

	// 消费所有产品
	for range ch {
	}

	// 检查通道是否关闭
	select {
	case _, ok := <-ch:
		if ok {
			t.Error("通道应该已关闭，但仍未关闭")
		}
	case <-time.After(1 * time.Second):
		t.Error("等待通道关闭超时")
	}
}
