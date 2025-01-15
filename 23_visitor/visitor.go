package visitor

import "fmt"

// Animal 接口，定义了接受访问者的方法
type Animal interface {
	Accept(visitor Visitor)
}

// Tiger 表示老虎
type Tiger struct{}

func (t *Tiger) Accept(visitor Visitor) {
	visitor.VisitTiger(t)
}

// Monkey 表示猴子
type Monkey struct{}

func (m *Monkey) Accept(visitor Visitor) {
	visitor.VisitMonkey(m)
}

// Elephant 表示大象
type Elephant struct{}

func (e *Elephant) Accept(visitor Visitor) {
	visitor.VisitElephant(e)
}

// Visitor 接口，定义了访问不同动物的行为
type Visitor interface {
	VisitTiger(tiger *Tiger)
	VisitMonkey(monkey *Monkey)
	VisitElephant(elephant *Elephant)
}

// Veterinarian 表示兽医
type Veterinarian struct{}

func (v *Veterinarian) VisitTiger(tiger *Tiger) {
	fmt.Println("兽医检查老虎的健康状况。")
}

func (v *Veterinarian) VisitMonkey(monkey *Monkey) {
	fmt.Println("兽医检查猴子的健康状况。")
}

func (v *Veterinarian) VisitElephant(elephant *Elephant) {
	fmt.Println("兽医检查大象的健康状况。")
}

// Zookeeper 表示饲养员
type Zookeeper struct{}

func (z *Zookeeper) VisitTiger(tiger *Tiger) {
	fmt.Println("饲养员给老虎喂肉。")
}

func (z *Zookeeper) VisitMonkey(monkey *Monkey) {
	fmt.Println("饲养员给猴子喂香蕉。")
}

func (z *Zookeeper) VisitElephant(elephant *Elephant) {
	fmt.Println("饲养员给大象喂草。")
}
