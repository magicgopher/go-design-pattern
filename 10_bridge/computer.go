package bridge

// Computer 接口
type Computer interface {
	Boot() error
	UseKeyboard() error
}

// MacBook 结构体
type MacBook struct {
	Keyboard Keyboard
}

// Boot 实现
func (m *MacBook) Boot() error {
	println("MacBook is booting up.")
	return nil
}

// UseKeyboard 实现
func (m *MacBook) UseKeyboard() error {
	print("Using ")
	if err := m.Keyboard.Type(); err != nil {
		return err
	}
	return nil
}

// GalaxyBook 结构体
type GalaxyBook struct {
	Keyboard Keyboard
}

// Boot 实现
func (g *GalaxyBook) Boot() error {
	println("GalaxyBook is booting up.")
	return nil
}

// UseKeyboard 实现
func (g *GalaxyBook) UseKeyboard() error {
	print("Using ")
	if err := g.Keyboard.Type(); err != nil {
		return err
	}
	return nil
}
