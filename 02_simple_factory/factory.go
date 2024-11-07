package simplefactory

// CoffeeFactory 是一个简单工厂，用于创建咖啡
type CoffeeFactory struct{}

// CreateCoffee 根据类型创建不同的咖啡
func (cf CoffeeFactory) CreateCoffee(coffeeType string) Coffee {
	switch coffeeType {
	case "latte":
		return Latte{}
	case "espresso":
		return Espresso{}
	default:
		return nil
	}
}
