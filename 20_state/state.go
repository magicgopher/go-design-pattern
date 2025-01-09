package state

import "fmt"

// State 接口定义了自动售货机的各种状态下的行为
type State interface {
	InsertCoin() error    // 投入硬币
	SelectProduct() error // 选择产品
	Dispense() error      // 发放产品
}

// NoProductState 结构体表示售货机没有产品的状态
type NoProductState struct {
	machine *VendingMachine
}

func (s *NoProductState) InsertCoin() error {
	return fmt.Errorf("无法接受硬币,产品已售罄")
}

func (s *NoProductState) SelectProduct() error {
	return fmt.Errorf("无法选择产品,产品已售罄")
}

func (s *NoProductState) Dispense() error {
	return fmt.Errorf("无法发放产品,产品已售罄")
}

// NoCoinState 结构体表示售货机有产品但没有硬币的状态
type NoCoinState struct {
	machine *VendingMachine
}

func (s *NoCoinState) InsertCoin() error {
	s.machine.hasCoin = true
	s.machine.setState()
	return nil
}

func (s *NoCoinState) SelectProduct() error {
	return fmt.Errorf("请先插入硬币")
}

func (s *NoCoinState) Dispense() error {
	return fmt.Errorf("请先插入硬币并选择产品")
}

// HasProductAndCoinState 结构体表示售货机有产品且已插入硬币的状态
type HasProductAndCoinState struct {
	machine *VendingMachine
}

func (s *HasProductAndCoinState) InsertCoin() error {
	return fmt.Errorf("已经插入了硬币,请选择产品或取回硬币")
}

func (s *HasProductAndCoinState) SelectProduct() error {
	fmt.Println("产品已选择,准备发放")
	return nil
}

func (s *HasProductAndCoinState) Dispense() error {
	fmt.Println("产品已发放,谢谢惠顾")
	return nil
}
