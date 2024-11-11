package builder_test

import (
	builder "github.com/magicgopher/go-design-pattern/06_builder"
	"testing"
)

// TestCarBuilder 建造者模式单元测试
func TestCarBuilder(t *testing.T) {
	// 创建一个具体的汽车构建器
	cb := builder.NewConcreteCarBuilder()
	// 创建一个导向器，使用构建器来构建汽车
	director := builder.NewDirector(cb)

	// 定义汽车的引擎
	engine := builder.Engine{Type: "V6", Power: 300}
	// 定义汽车的车身骨架
	skeleton := builder.Skeleton{Material: "Steel", Weight: 1200}
	// 定义汽车的车轮
	wheels := []builder.Wheel{
		{Type: "Alloy", Size: 18},
		{Type: "Alloy", Size: 18},
		{Type: "Alloy", Size: 18},
		{Type: "Alloy", Size: 18},
	}

	// 使用导向器构建具体的汽车对象
	car := director.Construct("Toyota", "Camry", 2023, "Red", engine, skeleton, wheels)

	// 验证汽车的制造商是否正确
	if car.Make != "Toyota" {
		t.Errorf("Expected Make to be 'Toyota', got '%s'", car.Make)
	}

	// 验证汽车的型号是否正确
	if car.Model != "Camry" {
		t.Errorf("Expected Model to be 'Camry', got '%s'", car.Model)
	}

	// 验证汽车的年份是否正确
	if car.Year != 2023 {
		t.Errorf("Expected Year to be 2023, got %d", car.Year)
	}

	// 验证汽车的颜色是否正确
	if car.Color != "Red" {
		t.Errorf("Expected Color to be 'Red', got '%s'", car.Color)
	}

	// 验证汽车的引擎类型是否正确
	if car.Engine.Type != "V6" {
		t.Errorf("Expected Engine Type to be 'V6', got '%s'", car.Engine.Type)
	}

	// 验证汽车的引擎马力是否正确
	if car.Engine.Power != 300 {
		t.Errorf("Expected Engine Power to be 300, got %d", car.Engine.Power)
	}

	// 验证汽车的车身骨架材料是否正确
	if car.Skeleton.Material != "Steel" {
		t.Errorf("Expected Skeleton Material to be 'Steel', got '%s'", car.Skeleton.Material)
	}

	// 验证汽车的车轮数量是否正确
	if len(car.Wheels) != 4 {
		t.Errorf("Expected 4 wheels, got %d", len(car.Wheels))
	}

	// 验证车轮的大小是否正确
	if car.Wheels[0].Size != 18 {
		t.Errorf("Expected Wheel Size to be 18, got %d", car.Wheels[0].Size)
	}
}
