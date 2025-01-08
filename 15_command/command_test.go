package command_test

import (
	command "github.com/magicgopher/go-design-pattern/15_command"
	"testing"
)

// TestRemoteControl
func TestRemoteControl(t *testing.T) {
	// 创建遥控器
	remote := command.NewRemoteControl()

	// 创建设备
	livingRoomLight := command.NewLight("Living Room")
	bedroomLight := command.NewLight("Bedroom")
	livingRoomTV := command.NewTV("Living Room")

	// 创建命令
	livingRoomLightOn := command.NewLightOnCommand(livingRoomLight)
	livingRoomLightOff := command.NewLightOffCommand(livingRoomLight)
	bedroomLightOn := command.NewLightOnCommand(bedroomLight)
	bedroomLightOff := command.NewLightOffCommand(bedroomLight)
	tvOn := command.NewTVOnCommand(livingRoomTV)
	tvOff := command.NewTVOffCommand(livingRoomTV)

	// 设置命令到遥控器
	remote.AddCommand(livingRoomLightOn)  // slot 0
	remote.AddCommand(livingRoomLightOff) // slot 1
	remote.AddCommand(bedroomLightOn)     // slot 2
	remote.AddCommand(bedroomLightOff)    // slot 3
	remote.AddCommand(tvOn)               // slot 4
	remote.AddCommand(tvOff)              // slot 5

	// 测试用例
	tests := []struct {
		name     string
		action   func() string
		expected string
	}{
		{
			name: "打开客厅灯",
			action: func() string {
				return remote.ExecuteCommand(0)
			},
			expected: "Living Room light is on",
		},
		{
			name: "关闭客厅灯",
			action: func() string {
				return remote.ExecuteCommand(1)
			},
			expected: "Living Room light is off",
		},
		{
			name: "打开卧室灯",
			action: func() string {
				return remote.ExecuteCommand(2)
			},
			expected: "Bedroom light is on",
		},
		{
			name: "打开电视",
			action: func() string {
				return remote.ExecuteCommand(4)
			},
			expected: "Living Room TV is on, channel 1",
		},
		{
			name: "撤销上一个命令(关闭电视)",
			action: func() string {
				return remote.UndoLastCommand()
			},
			expected: "Living Room TV is off",
		},
		{
			name: "执行无效命令",
			action: func() string {
				return remote.ExecuteCommand(10)
			},
			expected: "Invalid command slot",
		},
	}

	// 运行测试用例
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.action()
			if result != tt.expected {
				t.Errorf("期望得到 %v，但得到 %v", tt.expected, result)
			}
		})
	}
}

// TestMacroCommand
func TestMacroCommand(t *testing.T) {
	// 创建遥控器
	remote := command.NewRemoteControl()

	// 创建设备
	livingRoomLight := command.NewLight("Living Room")
	bedroomLight := command.NewLight("Bedroom")
	tv := command.NewTV("Living Room")

	// 创建回家模式的命令序列
	remote.AddCommand(command.NewLightOnCommand(livingRoomLight)) // slot 0
	remote.AddCommand(command.NewLightOnCommand(bedroomLight))    // slot 1
	remote.AddCommand(command.NewTVOnCommand(tv))                 // slot 2

	// 执行回家模式
	results := []string{
		remote.ExecuteCommand(0),
		remote.ExecuteCommand(1),
		remote.ExecuteCommand(2),
	}

	// 验证执行结果
	expectedResults := []string{
		"Living Room light is on",
		"Bedroom light is on",
		"Living Room TV is on, channel 1",
	}

	for i, result := range results {
		if result != expectedResults[i] {
			t.Errorf("回家模式测试失败: 期望得到 %v，但得到 %v", expectedResults[i], result)
		}
	}

	// 测试撤销操作
	for i := 0; i < 3; i++ {
		remote.UndoLastCommand()
	}

	// 验证设备状态
	if livingRoomLight.IsOn() {
		t.Error("客厅灯应该是关闭状态")
	}
}
