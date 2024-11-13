package static_proxy_test

import (
	staticproxy "github.com/magicgopher/go-design-pattern/07_proxy/static_proxy"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestStaticProxy 静态代理单元测试
func TestStaticProxy(t *testing.T) {
	// 创建实际的售卖者和代理商
	realSeller := &staticproxy.RealComputerSeller{}
	agent := &staticproxy.Agent{RealSeller: realSeller}

	// 代理商进行销售
	finalPrice := agent.Sell()

	// 断言客户支付的价格
	assert.Equal(t, 10798.8, finalPrice, "代理商扣除20%的利润后应支付10798.8")
}
