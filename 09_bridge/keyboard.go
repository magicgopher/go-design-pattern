package bridge

import "fmt"

// 键盘

// Keyboard 键盘
type Keyboard interface {
	// Type 键盘类型
	Type() error
}

// LogitechKeyboard 是具体的键盘实现（罗技键盘）
type LogitechKeyboard struct{}

func (l *LogitechKeyboard) Type() error {
	fmt.Println("Logitech Keyboard.")
	return nil
}

// RazerKeyboard 是具体的键盘实现（雷蛇键盘）
type RazerKeyboard struct{}

func (r *RazerKeyboard) Type() error {
	fmt.Println("Razer Keyboard.")
	return nil
}
