package simplefactory

// Coffee 是一个接口，表示咖啡
type Coffee interface {
	GetName() string
}

// Latte 是拿铁咖啡的具体实现
type Latte struct{}

func (l Latte) GetName() string {
	return "Latte"
}

// Espresso 是浓缩咖啡的具体实现
type Espresso struct{}

func (e Espresso) GetName() string {
	return "Espresso"
}
