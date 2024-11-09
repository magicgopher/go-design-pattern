package factorymethod

// 工厂方法

// PhoneFactory 手机工厂接口
type PhoneFactory interface {
	CreatePhone(name, color string) Phone
}

// 具体工厂

// ApplePhoneFactory Apple手机工厂
type ApplePhoneFactory struct {
}

func (apf *ApplePhoneFactory) CreatePhone(name, color string) Phone {
	return &IPhone{
		name:  name,
		color: color,
	}
}

// GalaxyFactory Samsung 工厂
type GalaxyFactory struct{}

func (g *GalaxyFactory) CreatePhone(name, color string) Phone {
	return &Galaxy{
		name:  name,
		color: color,
	}
}
