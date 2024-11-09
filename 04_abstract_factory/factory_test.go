package abstractfactory_test

import (
	abstractfactory "github.com/magicgopher/go-design-pattern/04_abstract_factory"
	"testing"
)

// TestAbstractFactory 测试抽象工厂模式
func TestAbstractFactory(t *testing.T) {
	// 创建苹果工厂
	appleFactory := &abstractfactory.AppleFactory{}
	applePhone := appleFactory.CreatePhone()
	appleLaptop := appleFactory.CreateLaptop()

	if applePhone.ProductMsg() != "iPhone 16 Pro Max" {
		t.Errorf("Expected iPhone 16 Pro Max, got %s", applePhone.ProductMsg())
	}
	if appleLaptop.ProductMsg() != "MacBook Pro 2024款 M4 Max 芯片" {
		t.Errorf("Expected MacBook Pro 2024款 M4 Max 芯片, got %s", appleLaptop.ProductMsg())
	}

	// 创建三星工厂
	samsungFactory := &abstractfactory.SamsungFactory{}
	samsungPhone := samsungFactory.CreatePhone()
	samsungLaptop := samsungFactory.CreateLaptop()

	if samsungPhone.ProductMsg() != "Galaxy S24 Ultra" {
		t.Errorf("Expected Galaxy S24 Ultra, got %s", samsungPhone.ProductMsg())
	}
	if samsungLaptop.ProductMsg() != "GalaxyBook4 Ultra" {
		t.Errorf("Expected GalaxyBook4 Ultra, got %s", samsungLaptop.ProductMsg())
	}
}
