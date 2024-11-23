package bridge_test

import (
	bridge "github.com/magicgopher/go-design-pattern/10_bridge"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestMacBookWithLogitech 测试 MacBook 使用 Logitech 键盘
func TestMacBookWithLogitech(t *testing.T) {
	// 创建 Logitech 键盘
	logitech := &bridge.LogitechKeyboard{}

	// 创建 MacBook
	macbook := &bridge.MacBook{Keyboard: logitech}

	// 启动 MacBook
	err := macbook.Boot()
	assert.Nil(t, err)

	// 使用 Logitech 键盘
	err = macbook.UseKeyboard()
	assert.Nil(t, err) // 期望使用 Logitech 键盘时没有错误
}

// TestGalaxyBookWithMustang 测试 GalaxyBook 使用 Mustang 键盘
func TestGalaxyBookWithMustang(t *testing.T) {
	// 创建 Razer 键盘
	razer := &bridge.RazerKeyboard{}
	// 创建 GalaxyBook
	galaxybook := &bridge.GalaxyBook{Keyboard: razer}
	// 启动 GalaxyBook
	err := galaxybook.Boot()
	assert.Nil(t, err)
	// 使用 Razer 键盘
	err = galaxybook.UseKeyboard()
	assert.Nil(t, err) // 期望使用 Razer 键盘时没有错误
}
