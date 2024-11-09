package factorymethod

import "fmt"

// Phone 手机接口
type Phone interface {
	// Info 基本信息
	Info() string
}

// Galaxy Samsung
type Galaxy struct {
	name  string // 名称
	color string // 颜色
}

// IPhone Apple
type IPhone struct {
	name  string // 名称
	color string // 颜色
}

// Info 基本信息
func (g *Galaxy) Info() string {
	return fmt.Sprintf("name: %s, color: %s\n", g.name, g.color)
}

// Info 基本信息
func (i *IPhone) Info() string {
	return fmt.Sprintf("name: %s, color: %s\n", i.name, i.color)
}
