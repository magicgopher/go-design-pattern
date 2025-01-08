package command

// 设备

// Light 电灯结构体
type Light struct {
	Location string
	isOn     bool
}

// NewLight 创建新的电灯
func NewLight(location string) *Light {
	return &Light{Location: location}
}

// TurnOn 打开电灯
func (l *Light) TurnOn() string {
	l.isOn = true
	return l.Location + " light is on"
}

// TurnOff 关闭电灯
func (l *Light) TurnOff() string {
	l.isOn = false
	return l.Location + " light is off"
}

// IsOn 获取电灯状态
func (l *Light) IsOn() bool {
	return l.isOn
}

// TV 电视结构体
type TV struct {
	Location string
	isOn     bool
	channel  int
}

// NewTV 创建新的电视
func NewTV(location string) *TV {
	return &TV{
		Location: location,
		channel:  1,
	}
}

func (t *TV) TurnOn() string {
	t.isOn = true
	return t.Location + " TV is on, channel " + string(rune(t.channel+'0'))
}

func (t *TV) TurnOff() string {
	t.isOn = false
	return t.Location + " TV is off"
}

func (t *TV) IsOn() bool {
	return t.isOn
}

func (t *TV) SetChannel(channel int) string {
	t.channel = channel
	return t.Location + " TV channel set to " + string(rune(channel+'0'))
}
