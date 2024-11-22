package adepter

// Voltage 电压接口
type Voltage interface {
	// ProvideVoltage 提供电压方法
	ProvideVoltage() int
}
