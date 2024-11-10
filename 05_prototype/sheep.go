package prototype

import "fmt"

// 原型模式

// 定义需要克隆的结构体

// Sheep 羊
type Sheep struct {
	Name   string  //名字
	Weight float64 // 体重
}

// Clone 返回一个新的 Sheep，复制当前对象
func (s *Sheep) Clone() Prototype {
	return &Sheep{
		Name:   s.Name,
		Weight: s.Weight,
	}
}

// String 输出 Sheep 的信息
func (s *Sheep) String() string {
	return fmt.Sprintf("Sheep{Name: %s, Weight: %0.2f}\n", s.Name, s.Weight)
}
