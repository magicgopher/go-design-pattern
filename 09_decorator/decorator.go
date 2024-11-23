package decorator

// MilkDecorator 牛奶的装饰器
type MilkDecorator struct {
	coffee Coffee
}

// NewMilkDecorator 返回一个 MilkDecorator
func NewMilkDecorator(c Coffee) *MilkDecorator {
	return &MilkDecorator{coffee: c}
}

// Cost 返回增加牛奶的咖啡的价格
func (m *MilkDecorator) Cost() float64 {
	return m.coffee.Cost() + 1.5 // 加牛奶的咖啡价格
}

// Description 返回增加牛奶的咖啡的描述
func (m *MilkDecorator) Description() string {
	return m.coffee.Description() + " + Milk"
}

// SugarDecorator 装饰器为咖啡添加糖
type SugarDecorator struct {
	coffee Coffee
}

// NewSugarDecorator 返回一个 SugarDecorator
func NewSugarDecorator(c Coffee) *SugarDecorator {
	return &SugarDecorator{coffee: c}
}

// Cost 返回增加糖的咖啡的价格
func (s *SugarDecorator) Cost() float64 {
	return s.coffee.Cost() + 0.5 // 加糖的咖啡价格
}

// Description 返回增加糖的咖啡的描述
func (s *SugarDecorator) Description() string {
	return s.coffee.Description() + " + Sugar"
}
