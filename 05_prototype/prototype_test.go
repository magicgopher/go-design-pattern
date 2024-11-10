package prototype_test

import (
	prototype "github.com/magicgopher/go-design-pattern/05_prototype"
	"reflect"
	"testing"
)

// TestSheepClone 原型模式单元测试
func TestSheepClone(t *testing.T) {
	// 创建一个 Sheep 对象
	original := &prototype.Sheep{Name: "Dolly", Weight: 45.5}

	// 克隆对象
	clone := original.Clone()

	// 检查克隆对象的类型和内容
	if reflect.TypeOf(clone) != reflect.TypeOf(original) {
		t.Errorf("Expected type %v, got %v", reflect.TypeOf(original), reflect.TypeOf(clone))
	}

	// 验证克隆对象的属性是否与原始对象一致
	if clone.(*prototype.Sheep).Name != original.Name || clone.(*prototype.Sheep).Weight != original.Weight {
		t.Errorf("Expected name %s and weight %.2f, got name %s and weight %.2f",
			original.Name, original.Weight,
			clone.(*prototype.Sheep).Name, clone.(*prototype.Sheep).Weight)
	}

	// 验证克隆和原始对象是不同的实例
	if original == clone {
		t.Errorf("Expected different instances for original and clone")
	}
}
