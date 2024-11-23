package decorator_test

import (
	decorator "github.com/magicgopher/go-design-pattern/09_decorator"
	"testing"
)

// TestSimpleCoffee 简单咖啡
func TestSimpleCoffee(t *testing.T) {
	coffee := &decorator.SimpleCoffee{}
	if coffee.Cost() != 5.0 {
		t.Errorf("expected 5.0, got %f", coffee.Cost())
	}
	if coffee.Description() != "Simple Coffee" {
		t.Errorf("expected 'Simple Coffee', got '%s'", coffee.Description())
	}
}

// TestMilkDecorator 增加牛奶
func TestMilkDecorator(t *testing.T) {
	coffee := &decorator.SimpleCoffee{}
	milkCoffee := decorator.NewMilkDecorator(coffee)

	if milkCoffee.Cost() != 6.5 {
		t.Errorf("expected 6.5, got %f", milkCoffee.Cost())
	}
	if milkCoffee.Description() != "Simple Coffee + Milk" {
		t.Errorf("expected 'Simple Coffee + Milk', got '%s'", milkCoffee.Description())
	}
}

// TestSugarDecorator 增加糖
func TestSugarDecorator(t *testing.T) {
	coffee := &decorator.SimpleCoffee{}
	sugarCoffee := decorator.NewSugarDecorator(coffee)

	if sugarCoffee.Cost() != 5.5 {
		t.Errorf("expected 5.5, got %f", sugarCoffee.Cost())
	}
	if sugarCoffee.Description() != "Simple Coffee + Sugar" {
		t.Errorf("expected 'Simple Coffee + Sugar', got '%s'", sugarCoffee.Description())
	}
}

// TestMultipleDecorators 多层装饰
func TestMultipleDecorators(t *testing.T) {
	coffee := &decorator.SimpleCoffee{}
	milkCoffee := decorator.NewMilkDecorator(coffee)
	sugarMilkCoffee := decorator.NewSugarDecorator(milkCoffee)

	if sugarMilkCoffee.Cost() != 7.0 {
		t.Errorf("expected 7.0, got %f", sugarMilkCoffee.Cost())
	}
	if sugarMilkCoffee.Description() != "Simple Coffee + Milk + Sugar" {
		t.Errorf("expected 'Simple Coffee + Milk + Sugar', got '%s'", sugarMilkCoffee.Description())
	}
}
