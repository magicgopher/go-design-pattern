package dynamicproxy_test

import (
	dynamicproxy "github.com/magicgopher/go-design-pattern/07_proxy/dynamic_proxy"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestDynamicProxy 动态代理单元测试
func TestDynamicProxy(t *testing.T) {
	// 创建实际的售卖者和代理商
	realSeller := &dynamicproxy.RealComputerSeller{}
	dynamicAgent := &dynamicproxy.DynamicAgent{RealSeller: realSeller}

	// 代理商进行销售
	finalPrice := dynamicAgent.Sell()

	// 断言客户支付的价格
	assert.Equal(t, 10798.8, finalPrice, "代理商扣除20%的利润后应支付10798.8")
}
