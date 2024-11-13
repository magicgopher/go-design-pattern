package static_proxy

import "fmt"

// 静态代理

// 案例：通过代理商购置电脑，代理商可以获取 20% 的利润

// ComputerSeller 接口，代表售卖电脑的功能
type ComputerSeller interface {
	Sell() float64 // 售卖
}

// RealComputerSeller 实现了ComputerSeller接口
type RealComputerSeller struct {
}

// Sell 方法实现
func (rcs *RealComputerSeller) Sell() float64 {
	// 电脑实际价格
	return 8999
}

// Agent 代理商，表示代理实际的销售
type Agent struct {
	RealSeller ComputerSeller
}

// Sell 方法实现，代理商从中获得20%的利润
func (a *Agent) Sell() float64 {
	originalPrice := a.RealSeller.Sell()
	commission := originalPrice * 0.20 // 20%的利润
	finalPrice := originalPrice + commission
	fmt.Printf("代理商销售电脑，原价: %.2f, 代理利润: %.2f, 最终销售价格: %.2f\n", originalPrice, commission, finalPrice)
	return finalPrice
}
