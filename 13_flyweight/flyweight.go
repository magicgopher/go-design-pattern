package flyweight

import "fmt"

// TreeType 表示树的共享部分（内部状态）
type TreeType struct {
	Name    string // 名称
	Color   string // 颜色
	Texture string // 纹理
}

// TreeTypeFactory 管理共享的 TreeType 对象
type TreeTypeFactory struct {
	treeTypes map[string]*TreeType
}

// NewTreeTypeFactory 创建 TreeTypeFactory 实例
func NewTreeTypeFactory() *TreeTypeFactory {
	return &TreeTypeFactory{treeTypes: make(map[string]*TreeType)}
}

// GetTreeType 获取或创建共享的 TreeType 对象
func (f *TreeTypeFactory) GetTreeType(name, color, texture string) *TreeType {
	key := name + "_" + color + "_" + texture
	if treeType, exists := f.treeTypes[key]; exists {
		return treeType
	}
	// 如果对象不存在，则创建并加入池中
	treeType := &TreeType{Name: name, Color: color, Texture: texture}
	f.treeTypes[key] = treeType
	return treeType
}

// Tree 表示游戏中的树对象（包含外部状态）
type Tree struct {
	X, Y     int
	TreeType *TreeType
}

// Draw 绘制树的方法
func (t *Tree) Draw() {
	fmt.Printf("Drawing tree '%s' at (%d, %d) with color '%s' and texture '%s'\n",
		t.TreeType.Name, t.X, t.Y, t.TreeType.Color, t.TreeType.Texture)
}

// Forest 表示森林（包含许多树）
type Forest struct {
	trees   []*Tree
	factory *TreeTypeFactory
}

// NewForest 创建 Forest 实例
func NewForest() *Forest {
	return &Forest{
		trees:   make([]*Tree, 0),
		factory: NewTreeTypeFactory(),
	}
}

// PlantTree 在森林中种植一棵树
func (f *Forest) PlantTree(x, y int, name, color, texture string) {
	treeType := f.factory.GetTreeType(name, color, texture)
	tree := &Tree{X: x, Y: y, TreeType: treeType}
	f.trees = append(f.trees, tree)
}

// Draw 绘制森林中的所有树
func (f *Forest) Draw() {
	for _, tree := range f.trees {
		tree.Draw()
	}
}

// GetTreeTypeCount 返回共享的 TreeType 对象数量
func (f *Forest) GetTreeTypeCount() int {
	return len(f.factory.treeTypes)
}
