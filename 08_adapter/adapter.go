package adepter

// HighVoltage 结构体表示220V电压
type HighVoltage struct{}

// ProvideVoltage 方法返回220V电压
func (hv HighVoltage) ProvideVoltage() int {
	return 220
}

// PowerAdapter 适配器，将220V转换为5V
type PowerAdapter struct {
	HighVoltage HighVoltage
}

// ProvideVoltage 方法返回5V电压
func (pa PowerAdapter) ProvideVoltage() int {
	return 5
}
