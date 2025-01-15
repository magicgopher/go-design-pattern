package visitor_test

import (
	visitor "github.com/magicgopher/go-design-pattern/23_visitor"
	"testing"
)

func TestVisitorPattern(t *testing.T) {
	// 创建动物
	tiger := &visitor.Tiger{}
	monkey := &visitor.Monkey{}
	elephant := &visitor.Elephant{}

	// 创建访问者
	veterinarian := &visitor.Veterinarian{}
	zookeeper := &visitor.Zookeeper{}

	// 兽医访问动物
	t.Log("兽医检查动物：")
	tiger.Accept(veterinarian)
	monkey.Accept(veterinarian)
	elephant.Accept(veterinarian)

	// 饲养员访问动物
	t.Log("\n饲养员喂养动物：")
	tiger.Accept(zookeeper)
	monkey.Accept(zookeeper)
	elephant.Accept(zookeeper)
}
