package prototype

import "fmt"

// 原型模式

// 定义需要克隆的结构体

// Sheep 羊
type Sheep struct {
	name   string  //名字
	weight float64 // 体重
}

// Clone 返回一个新的 Sheep，复制当前对象
func (s *Sheep) Clone() Prototype {
	return &Sheep{
		name:   s.name,
		weight: s.weight,
	}
}

// String 输出 Sheep 的信息
func (s *Sheep) String() string {
	return fmt.Sprintf("Sheep{Name: %s, Weight: %0.2f}\n", s.name, s.weight)
}
