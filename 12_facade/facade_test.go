package facade_test

import (
	facade "github.com/magicgopher/go-design-pattern/12_facade"
	"testing"
)

// TestHomeTheaterFacade 测试 HomeTheaterFacade 的功能
func TestHomeTheaterFacade(t *testing.T) {
	// 创建家庭影院设备
	tv := &facade.TV{}
	// 创建音响系统
	soundSystem := &facade.SoundSystem{}
	// 创建 DVD 播放器
	dvdPlayer := &facade.DVDPlayer{}
	// 创建 HomeTheaterFacade（家庭影院设备的外观）
	homeTheater := facade.NewHomeTheaterFacade(tv, soundSystem, dvdPlayer)

	// 准备观看电影
	homeTheater.WatchMovie()
	// 关闭家庭影院设备
	homeTheater.EndMovie()
}
