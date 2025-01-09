package state

// VendingMachine 自动售货机
type VendingMachine struct {
	hasProduct             bool  // 表示售货机是否有产品
	hasCoin                bool  // 表示是否已插入硬币
	currentState           State // 当前售货机的状态
	noProductState         State // 无产品状态
	noCoinState            State // 无硬币状态
	hasProductAndCoinState State // 有产品且已插入硬币的状态
}

// NewVendingMachine 创建一个新的自动售货机实例
func NewVendingMachine(productCount int) *VendingMachine {
	v := &VendingMachine{
		hasProduct: productCount > 0,
		hasCoin:    false,
	}
	v.noProductState = &NoProductState{machine: v}
	v.noCoinState = &NoCoinState{machine: v}
	v.hasProductAndCoinState = &HasProductAndCoinState{machine: v}
	v.setState()
	return v
}

// setState 根据当前条件设置售货机的状态
func (v *VendingMachine) setState() {
	if !v.hasProduct {
		v.currentState = v.noProductState
	} else if !v.hasCoin {
		v.currentState = v.noCoinState
	} else {
		v.currentState = v.hasProductAndCoinState
	}
}

// InsertCoin 插入硬币
func (v *VendingMachine) InsertCoin() error {
	return v.currentState.InsertCoin()
}

// SelectProduct 选择产品
func (v *VendingMachine) SelectProduct() error {
	return v.currentState.SelectProduct()
}

// Dispense 发放产品
func (v *VendingMachine) Dispense() error {
	err := v.currentState.Dispense()
	if err == nil {
		v.hasProduct = false
		v.hasCoin = false
		v.setState()
	}
	return err
}
