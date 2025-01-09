package memento

import "fmt"

// 备忘录模式

// Memento 结构体用于存储编辑器的状态
type Memento struct {
	content string
}

// GetContent 方法用于获取备忘录中存储的内容
func (m *Memento) GetContent() string {
	return m.content
}

// Editor 结构体代表文本编辑器
type Editor struct {
	content string
}

// NewEditor 创建一个新的编辑器实例
func NewEditor() *Editor {
	return &Editor{}
}

// Type 方法用于在编辑器中输入内容
func (e *Editor) Type(text string) {
	e.content += text
}

// GetContent 方法用于获取编辑器当前的内容
func (e *Editor) GetContent() string {
	return e.content
}

// Save 方法用于保存编辑器的当前状态
func (e *Editor) Save() *Memento {
	return &Memento{content: e.content}
}

// Restore 方法用于从备忘录中恢复编辑器的状态
func (e *Editor) Restore(m *Memento) {
	e.content = m.GetContent()
}

// CareTaker 结构体用于管理备忘录
type CareTaker struct {
	mementos []*Memento
}

// NewCareTaker 创建一个新的管理者实例
func NewCareTaker() *CareTaker {
	return &CareTaker{mementos: make([]*Memento, 0)}
}

// AddMemento 方法用于添加一个新的备忘录
func (c *CareTaker) AddMemento(m *Memento) {
	c.mementos = append(c.mementos, m)
}

// GetMemento 方法用于获取指定索引的备忘录
func (c *CareTaker) GetMemento(index int) (*Memento, error) {
	if index < 0 || index >= len(c.mementos) {
		return nil, fmt.Errorf("invalid index")
	}
	return c.mementos[index], nil
}
