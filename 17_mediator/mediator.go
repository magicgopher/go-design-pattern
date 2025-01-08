package mediator

import "fmt"

// 中介者模式

// Mediator 接口定义了中介者的行为
type Mediator interface {
	SendMessage(message string, user User)
	RegisterUser(user User)
	BroadcastMessage(message string, sender User)
	SendPrivateMessage(message string, sender User, receiver User)
}

// ChatRoom 是具体的中介者实现
type ChatRoom struct {
	users []User
}

func NewChatRoom() *ChatRoom {
	return &ChatRoom{}
}

func (c *ChatRoom) SendMessage(message string, user User) {
	fmt.Printf("[%s] %s\n", user.GetName(), message)
}

func (c *ChatRoom) RegisterUser(user User) {
	c.users = append(c.users, user)
	c.SendMessage("已加入聊天室", user)
}

func (c *ChatRoom) BroadcastMessage(message string, sender User) {
	for _, user := range c.users {
		if user != sender {
			c.SendMessage(fmt.Sprintf("广播消息来自 %s: %s", sender.GetName(), message), user)
		}
	}
}

func (c *ChatRoom) SendPrivateMessage(message string, sender User, receiver User) {
	c.SendMessage(fmt.Sprintf("私信来自 %s: %s", sender.GetName(), message), receiver)
}

// User 接口定义了用户的基本行为
type User interface {
	SetMediator(m Mediator)
	GetName() string
	Send(message string)
	SendPrivate(message string, receiver User)
	Receive(message string)
}

// BaseUser 包含所有用户共有的属性和方法
type BaseUser struct {
	name     string
	mediator Mediator
}

func (u *BaseUser) SetMediator(m Mediator) {
	u.mediator = m
}

func (u *BaseUser) GetName() string {
	return u.name
}

func (u *BaseUser) Send(message string) {
	u.mediator.BroadcastMessage(message, u)
}

func (u *BaseUser) SendPrivate(message string, receiver User) {
	u.mediator.SendPrivateMessage(message, u, receiver)
}

func (u *BaseUser) Receive(message string) {
	fmt.Printf("%s 收到消息: %s\n", u.name, message)
}

// Admin 代表管理员用户
type Admin struct {
	BaseUser
}

func NewAdmin(name string) *Admin {
	return &Admin{BaseUser{name: name}}
}

func (a *Admin) SendAnnouncement(message string) {
	fmt.Printf("系统公告: %s\n", message)
	a.mediator.BroadcastMessage(fmt.Sprintf("系统公告: %s", message), a)
}

// Member 代表普通会员
type Member struct {
	BaseUser
}

func NewMember(name string) *Member {
	return &Member{BaseUser{name: name}}
}

// Guest 代表访客用户
type Guest struct {
	BaseUser
}

func NewGuest(name string) *Guest {
	return &Guest{BaseUser{name: name}}
}

func (g *Guest) Send(message string) {
	fmt.Printf("访客 %s 无法发送消息\n", g.name)
}

func (g *Guest) SendPrivate(message string, receiver User) {
	fmt.Printf("访客 %s 无法发送私信\n", g.name)
}
