package facade

import "fmt"

// TV 代表一台电视
type TV struct{}

// On 打开电视
func (t *TV) On() {
	fmt.Println("打开电视.")
}

// Off 关闭电视
func (t *TV) Off() {
	fmt.Println("关闭电视.")
}

// SoundSystem 代表一个音响系统
type SoundSystem struct{}

// On 打开音响系统
func (s *SoundSystem) On() {
	fmt.Println("打开音响系统.")
}

// Off 关闭音响系统
func (s *SoundSystem) Off() {
	fmt.Println("关闭音响系统.")
}

// DVDPlayer 代表一个 DVD 播放器
type DVDPlayer struct{}

// InsertDisc 插入一张光盘
func (d *DVDPlayer) InsertDisc() {
	fmt.Println("插入一张光盘.")
}

// Play 播放 DVD
func (d *DVDPlayer) Play() {
	fmt.Println("播放 DVD.")
}

// Stop 停止播放 DVD
func (d *DVDPlayer) Stop() {
	fmt.Println("停止播放 DVD.")
}

// HomeTheaterFacade 为家庭影院设备提供一个统一的接口
type HomeTheaterFacade struct {
	tv          *TV
	soundSystem *SoundSystem
	dvdPlayer   *DVDPlayer
}

// NewHomeTheaterFacade 创建一个新的 HomeTheaterFacade 实例
func NewHomeTheaterFacade(tv *TV, soundSystem *SoundSystem, dvdPlayer *DVDPlayer) *HomeTheaterFacade {
	return &HomeTheaterFacade{
		tv:          tv,
		soundSystem: soundSystem,
		dvdPlayer:   dvdPlayer,
	}
}

// WatchMovie 准备观看电影
func (h *HomeTheaterFacade) WatchMovie() {
	fmt.Println("准备观看电影...")
	h.tv.On()
	h.soundSystem.On()
	h.dvdPlayer.InsertDisc()
	h.dvdPlayer.Play()
	fmt.Println("享受电影吧！")
}

// EndMovie 关闭家庭影院设备
func (h *HomeTheaterFacade) EndMovie() {
	fmt.Println("关闭影音设备...")
	h.dvdPlayer.Stop()
	h.soundSystem.Off()
	h.tv.Off()
}
