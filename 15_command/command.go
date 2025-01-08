package command

// Command 命令接口
type Command interface {
	Execute() string
	Undo() string
}

// LightOnCommand 开灯命令
type LightOnCommand struct {
	light *Light
}

func NewLightOnCommand(light *Light) *LightOnCommand {
	return &LightOnCommand{light: light}
}

func (c *LightOnCommand) Execute() string {
	return c.light.TurnOn()
}

func (c *LightOnCommand) Undo() string {
	return c.light.TurnOff()
}

// LightOffCommand 关灯命令
type LightOffCommand struct {
	light *Light
}

func NewLightOffCommand(light *Light) *LightOffCommand {
	return &LightOffCommand{light: light}
}

func (c *LightOffCommand) Execute() string {
	return c.light.TurnOff()
}

func (c *LightOffCommand) Undo() string {
	return c.light.TurnOn()
}

// TVOnCommand 打开电视命令
type TVOnCommand struct {
	tv *TV
}

func NewTVOnCommand(tv *TV) *TVOnCommand {
	return &TVOnCommand{tv: tv}
}

func (c *TVOnCommand) Execute() string {
	return c.tv.TurnOn()
}

func (c *TVOnCommand) Undo() string {
	return c.tv.TurnOff()
}

// TVOffCommand 关闭电视命令
type TVOffCommand struct {
	tv *TV
}

func NewTVOffCommand(tv *TV) *TVOffCommand {
	return &TVOffCommand{tv: tv}
}

func (c *TVOffCommand) Execute() string {
	return c.tv.TurnOff()
}

func (c *TVOffCommand) Undo() string {
	return c.tv.TurnOn()
}

// RemoteControl 遥控器
type RemoteControl struct {
	commands []Command
	history  []Command
}

// NewRemoteControl 创建新的遥控器
func NewRemoteControl() *RemoteControl {
	return &RemoteControl{
		commands: make([]Command, 0),
		history:  make([]Command, 0),
	}
}

// AddCommand 添加命令
func (r *RemoteControl) AddCommand(cmd Command) {
	r.commands = append(r.commands, cmd)
}

// ExecuteCommand 执行指定位置的命令
func (r *RemoteControl) ExecuteCommand(slot int) string {
	if slot >= 0 && slot < len(r.commands) {
		cmd := r.commands[slot]
		r.history = append(r.history, cmd)
		return cmd.Execute()
	}
	return "Invalid command slot"
}

// UndoLastCommand 撤销上一个命令
func (r *RemoteControl) UndoLastCommand() string {
	if len(r.history) > 0 {
		lastCmd := r.history[len(r.history)-1]
		r.history = r.history[:len(r.history)-1]
		return lastCmd.Undo()
	}
	return "No commands to undo"
}
