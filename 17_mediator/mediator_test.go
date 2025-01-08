package mediator_test

import (
	mediator "github.com/magicgopher/go-design-pattern/17_mediator"
	"testing"
)

// 单元测试

// TestChatRoom 中介者模式单元测试
func TestChatRoom(t *testing.T) {
	chatRoom := mediator.NewChatRoom()

	admin := mediator.NewAdmin("管理员小张")
	member1 := mediator.NewMember("会员小李")
	member2 := mediator.NewMember("会员小王")
	guest := mediator.NewGuest("访客小明")

	chatRoom.RegisterUser(admin)
	chatRoom.RegisterUser(member1)
	chatRoom.RegisterUser(member2)
	chatRoom.RegisterUser(guest)

	admin.SetMediator(chatRoom)
	member1.SetMediator(chatRoom)
	member2.SetMediator(chatRoom)
	guest.SetMediator(chatRoom)

	t.Run("Test broadcast message", func(t *testing.T) {
		member1.Send("大家好!")
	})

	t.Run("Test private message", func(t *testing.T) {
		member1.SendPrivate("你好,小王!", member2)
	})

	t.Run("Test admin announcement", func(t *testing.T) {
		admin.SendAnnouncement("欢迎来到我们的聊天室!")
	})

	t.Run("Test guest restrictions", func(t *testing.T) {
		guest.Send("我可以说话吗?")
		guest.SendPrivate("私密消息", member1)
	})
}
