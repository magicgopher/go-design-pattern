package dynamicproxy

import (
	"fmt"
	"reflect"
)

// 动态代理
// 反射的的方式实现

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

// DynamicAgent 动态代理
type DynamicAgent struct {
	RealSeller ComputerSeller
}

// Sell 方法实现，使用反射来动态调用
func (da *DynamicAgent) Sell() float64 {
	sellerValue := reflect.ValueOf(da.RealSeller)
	sellerMethod := sellerValue.MethodByName("Sell")
	if sellerMethod.IsValid() {
		// 调用实际的售卖方法
		originalPrice := sellerMethod.Call(nil)[0].Interface().(float64)
		commission := originalPrice * 0.20 // 20%的利润
		finalPrice := originalPrice + commission
		fmt.Printf("代理商销售电脑，原价: %.2f, 代理利润: %.2f, 最终销售价格: %.2f\n", originalPrice, commission, finalPrice)
		return finalPrice
	}
	return 0
}
