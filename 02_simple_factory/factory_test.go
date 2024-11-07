package simplefactory_test

import (
	simplefactory "github.com/MagicGopher/go-design-pattern/02_simple_factory"
	"testing"
)

// TestCreateCoffee 测试咖啡工厂的创建功能
func TestCreateCoffee(t *testing.T) {
	// 创建一个咖啡工厂实例
	factory := simplefactory.CoffeeFactory{}

	// 测试创建拿铁咖啡
	latte := factory.CreateCoffee("latte")
	if latte == nil || latte.GetName() != "Latte" {
		t.Errorf("期望拿铁, 得到 %v", latte.GetName())
	}

	// 测试创建浓缩咖啡
	espresso := factory.CreateCoffee("espresso")
	if espresso == nil || espresso.GetName() != "Espresso" {
		t.Errorf("期望浓缩咖啡, 得到 %v", espresso.GetName())
	}

	// 测试创建未知类型的咖啡
	unknown := factory.CreateCoffee("unknown")
	if unknown != nil {
		t.Errorf("期望 nil, 得到 %v", unknown.GetName())
	}
}
