package abstractfactory

// ElectronAbstractProduct 抽象电子产品接口
type ElectronAbstractProduct interface {
	ProductMsg() string
}

// Apple 产品实现

// IPhone Apple 手机产品
type IPhone struct{}

func (i *IPhone) ProductMsg() string {
	return "iPhone 16 Pro Max"
}

// MacBook Apple 笔记本产品
type MacBook struct{}

func (m *MacBook) ProductMsg() string {
	return "MacBook Pro 2024款 M4 Max 芯片"
}

// Samsung 产品实现

// GalaxyPhone Samsung 手机产品
type GalaxyPhone struct{}

func (gp *GalaxyPhone) ProductMsg() string {
	return "Galaxy S24 Ultra"
}

// GalaxyBook Samsung 笔记本产品
type GalaxyBook struct{}

func (gb *GalaxyBook) ProductMsg() string {
	return "GalaxyBook4 Ultra"
}
