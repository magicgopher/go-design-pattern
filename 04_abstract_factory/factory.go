package abstractfactory

// AbstractFactory 抽象工厂接口
type AbstractFactory interface {
	CreatePhone() ElectronAbstractProduct
	CreateLaptop() ElectronAbstractProduct
}

// AppleFactory 具体的苹果工厂
type AppleFactory struct{}

func (af *AppleFactory) CreatePhone() ElectronAbstractProduct {
	return &IPhone{}
}

func (af *AppleFactory) CreateLaptop() ElectronAbstractProduct {
	return &MacBook{}
}

// SamsungFactory 具体的三星工厂
type SamsungFactory struct{}

func (sf *SamsungFactory) CreatePhone() ElectronAbstractProduct {
	return &GalaxyPhone{}
}

func (sf *SamsungFactory) CreateLaptop() ElectronAbstractProduct {
	return &GalaxyBook{}
}
