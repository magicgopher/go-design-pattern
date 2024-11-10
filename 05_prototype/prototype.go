package prototype

// Prototype 是一个接口，定义了克隆方法
type Prototype interface {
	Clone() Prototype
}
