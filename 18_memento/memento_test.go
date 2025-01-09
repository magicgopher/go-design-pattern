package memento_test

import (
	memento "github.com/magicgopher/go-design-pattern/18_memento"
	"testing"
)

// TestEditor 备忘录模式单元测试
func TestEditor(t *testing.T) {
	editor := memento.NewEditor()
	careTaker := memento.NewCareTaker()

	// 第一次输入
	editor.Type("Hello, ")
	careTaker.AddMemento(editor.Save())

	// 第二次输入
	editor.Type("World!")
	careTaker.AddMemento(editor.Save())

	// 检查当前内容
	if editor.GetContent() != "Hello, World!" {
		t.Errorf("Expected content to be 'Hello, World!', but got '%s'", editor.GetContent())
	}

	// 第三次输入
	editor.Type(" How are you?")

	// 检查当前内容
	if editor.GetContent() != "Hello, World! How are you?" {
		t.Errorf("Expected content to be 'Hello, World! How are you?', but got '%s'", editor.GetContent())
	}

	// 撤销到第二次输入后的状态
	if memento, err := careTaker.GetMemento(1); err == nil {
		editor.Restore(memento)
	} else {
		t.Errorf("Failed to get memento: %v", err)
	}

	// 检查恢复后的内容
	if editor.GetContent() != "Hello, World!" {
		t.Errorf("Expected content to be 'Hello, World!', but got '%s'", editor.GetContent())
	}

	// 撤销到第一次输入后的状态
	if memento, err := careTaker.GetMemento(0); err == nil {
		editor.Restore(memento)
	} else {
		t.Errorf("Failed to get memento: %v", err)
	}

	// 检查恢复后的内容
	if editor.GetContent() != "Hello, " {
		t.Errorf("Expected content to be 'Hello, ', but got '%s'", editor.GetContent())
	}
}
