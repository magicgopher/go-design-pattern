package decorator

// Coffee 咖啡接口
type Coffee interface {
	Description() string
	Cost() float64
}

// SimpleCoffee 是最基本的咖啡类型
type SimpleCoffee struct{}

// Cost SimpleCoffee 实现了 Coffee 接口
func (c *SimpleCoffee) Cost() float64 {
	return 5.0 // 基本咖啡的价格
}

// Description SimpleCoffee 实现了 Coffee 接口
func (c *SimpleCoffee) Description() string {
	return "Simple Coffee"
}
