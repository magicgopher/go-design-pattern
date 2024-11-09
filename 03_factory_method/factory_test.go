package factorymethod_test

import (
	factorymethod "github.com/magicgopher/go-design-pattern/03_factory_method"
	"testing"
)

// TestPhoneFactory 工厂方法单元测试
func TestPhoneFactory(t *testing.T) {
	// 测试 Galaxy 工厂，创建不同配置的 Galaxy
	galaxyFactory := &factorymethod.GalaxyFactory{}
	galaxyPhone1 := galaxyFactory.CreatePhone("Galaxy S22", "Phantom Black")
	galaxyPhone2 := galaxyFactory.CreatePhone("Galaxy Z Flip", "Cream")

	if galaxyPhone1.Info() != "name: Galaxy S22, color: Phantom Black\n" {
		t.Errorf("Expected Galaxy S22 with Phantom Black color, but got: %s", galaxyPhone1.Info())
	}
	if galaxyPhone2.Info() != "name: Galaxy Z Flip, color: Cream\n" {
		t.Errorf("Expected Galaxy Z Flip with Cream color, but got: %s", galaxyPhone2.Info())
	}

	// 测试 iPhone 工厂，创建不同配置的 iPhone
	iphoneFactory := &factorymethod.ApplePhoneFactory{}
	iphone1 := iphoneFactory.CreatePhone("iPhone 14", "Midnight")
	iphone2 := iphoneFactory.CreatePhone("iPhone 14 Pro", "Gold")

	if iphone1.Info() != "name: iPhone 14, color: Midnight\n" {
		t.Errorf("Expected iPhone 14 with Midnight color, but got: %s", iphone1.Info())
	}
	if iphone2.Info() != "name: iPhone 14 Pro, color: Gold\n" {
		t.Errorf("Expected iPhone 14 Pro with Gold color, but got: %s", iphone2.Info())
	}
}
