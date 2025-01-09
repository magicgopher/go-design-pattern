package state_test

import (
	state "github.com/magicgopher/go-design-pattern/20_state"
	"testing"
)

func TestVendingMachine(t *testing.T) {

	t.Run("测试有商品的自动售货机", func(t *testing.T) {

		vm := state.NewVendingMachine(1)

		if err := vm.SelectProduct(); err == nil {
			t.Error("选择商品时未投币应该报错,但是得到了nil")
		}

		if err := vm.InsertCoin(); err != nil {
			t.Errorf("投币时不应该报错,但是得到了错误: %v", err)
		}

		if err := vm.SelectProduct(); err != nil {
			t.Errorf("选择商品时不应该报错,但是得到了错误: %v", err)
		}

		if err := vm.Dispense(); err != nil {
			t.Errorf("出货时不应该报错,但是得到了错误: %v", err)
		}

		if err := vm.SelectProduct(); err == nil {
			t.Error("出货后选择商品应该报错,但是得到了nil")
		}
	})

	t.Run("测试无商品的自动售货机", func(t *testing.T) {

		vm := state.NewVendingMachine(0)

		if err := vm.InsertCoin(); err == nil {
			t.Error("无商品时投币应该报错,但是得到了nil")
		}

		if err := vm.SelectProduct(); err == nil {
			t.Error("无商品时选择商品应该报错,但是得到了nil")
		}

		if err := vm.Dispense(); err == nil {
			t.Error("无商品时出货应该报错,但是得到了nil")
		}
	})
}
