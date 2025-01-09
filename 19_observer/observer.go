package observer

import "fmt"

//

// Observer 观察者接口定义
type Observer interface {
	Update(string) // 更新
}

// Subject 主题
type Subject interface {
	Register(Observer)   // 注册
	Deregister(Observer) // 移除
	NotifyAll()          // 通知全部
}

// NewsPublisher 结构体实现了 Subject 接口
type NewsPublisher struct {
	observers []Observer
	news      string
}

// NewNewsPublisher 创建一个新的 NewsPublisher 实例
func NewNewsPublisher() *NewsPublisher {
	return &NewsPublisher{
		observers: make([]Observer, 0),
	}
}

// Register 方法用于注册观察者
func (np *NewsPublisher) Register(o Observer) {
	np.observers = append(np.observers, o)
}

// Deregister 方法用于移除观察者
func (np *NewsPublisher) Deregister(o Observer) {
	for i, obs := range np.observers {
		if obs == o {
			np.observers = append(np.observers[:i], np.observers[i+1:]...)
			break
		}
	}
}

// NotifyAll 方法用于通知所有观察者
func (np *NewsPublisher) NotifyAll() {
	for _, observer := range np.observers {
		observer.Update(np.news)
	}
}

// SetNews 方法用于设置新闻并通知所有观察者
func (np *NewsPublisher) SetNews(news string) {
	np.news = news
	np.NotifyAll()
}

// NewsReader 结构体实现了 Observer 接口
type NewsReader struct {
	id string
}

// NewNewsReader 创建一个新的 NewsReader 实例
func NewNewsReader(id string) *NewsReader {
	return &NewsReader{id: id}
}

// Update 方法用于接收并处理更新
func (nr *NewsReader) Update(news string) {
	fmt.Printf("Reader %s received news: %s\n", nr.id, news)
}
