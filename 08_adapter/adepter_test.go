package adepter_test

import (
	adapter "github.com/magicgopher/go-design-pattern/08_adapter"
	"testing"
)

// 适配器模式单元测试
func TestPowerAdapter(t *testing.T) {
	// 创建被适配的对象
	highVoltage := adapter.HighVoltage{}
	// 创建适配器
	powerAdapter := adapter.PowerAdapter{HighVoltage: highVoltage}

	// 验证适配器转换后的电压
	if got := powerAdapter.ProvideVoltage(); got != 5 {
		t.Errorf("ProvideVoltage() = %v; want 5", got)
	}
}
