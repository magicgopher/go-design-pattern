package observer_test

import (
	observer "github.com/magicgopher/go-design-pattern/19_observer"
	"testing"
)

// 单元测试

// TestNewsPublisher 观察者模式单元测试
func TestNewsPublisher(t *testing.T) {
	publisher := observer.NewNewsPublisher()

	reader1 := observer.NewNewsReader("1")
	reader2 := observer.NewNewsReader("2")
	reader3 := observer.NewNewsReader("3")

	publisher.Register(reader1)
	publisher.Register(reader2)
	publisher.Register(reader3)

	t.Run("Test notifying all observers", func(t *testing.T) {
		publisher.SetNews("Breaking news: Go 2.0 released!")
	})

	t.Run("Test deregistering an observer", func(t *testing.T) {
		publisher.Deregister(reader2)
		publisher.SetNews("Weather update: Sunny day ahead!")
	})

	t.Run("Test registering a new observer", func(t *testing.T) {
		reader4 := observer.NewNewsReader("4")
		publisher.Register(reader4)
		publisher.SetNews("Sports news: Local team wins championship!")
	})
}
