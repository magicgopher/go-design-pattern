package template

import "fmt"

// CarMaker 定义制作汽车的模版方法接口
type CarMaker interface {
	BuildChassis()  // 构建底盘
	InstallEngine() // 安装引擎
	Paint()         // 喷漆
	AssembleBody()  // 组装车身主体
	Complete()      // 完成制造
}

// CarTemplate 提供模版方法的实现
type CarTemplate struct {
	Car CarMaker
}

// MakeCar 模版方法，定义制作汽车的完整流程
func (c *CarTemplate) MakeCar() {
	fmt.Println("Starting car production...")
	c.Car.BuildChassis()
	c.Car.InstallEngine()
	c.Car.Paint()
	c.Car.AssembleBody()
	c.Car.Complete()
	fmt.Println("Car production completed.")
}
