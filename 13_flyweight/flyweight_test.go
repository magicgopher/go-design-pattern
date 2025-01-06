package flyweight_test

import (
	flyweight "github.com/magicgopher/go-design-pattern/13_flyweight"
	"testing"
)

func TestFlyweight(t *testing.T) {
	// 创建一个森林实例
	forest := flyweight.NewForest()

	// 在森林中种植一些树
	forest.PlantTree(1, 2, "橡树", "绿色", "粗糙的纹理")  // 第一棵树
	forest.PlantTree(3, 4, "松树", "深绿色", "光滑的纹理") // 第二棵树
	forest.PlantTree(5, 6, "橡树", "绿色", "粗糙的纹理")  // 共享 "橡树" 类型
	forest.PlantTree(7, 8, "松树", "深绿色", "光滑的纹理") // 共享 "松树" 类型

	// 绘制森林中的所有树
	forest.Draw()

	// 验证共享的 TreeType 对象的数量是否正确
	// 因为只有两种类型的树（橡树和松树），共享对象池中应有 2 个 TreeType 对象
	expectedTreeTypeCount := 2
	actualTreeTypeCount := forest.GetTreeTypeCount()

	if actualTreeTypeCount != expectedTreeTypeCount {
		t.Errorf("期望的 TreeType 对象数量为 %d，但实际得到的是 %d", expectedTreeTypeCount, actualTreeTypeCount)
	}
}
