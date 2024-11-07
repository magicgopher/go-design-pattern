package singleton

import "sync"

// 定义全局变量
var singleton *Singleton

// 定义once全局变量
var once sync.Once

// Singleton 定义一个单例结构体
type Singleton struct {
}

// GetInstance 获取Singleton结构体实例变量函数
func GetInstance() *Singleton {
	// 确保只初始化一次Singleton实例
	once.Do(func() {
		singleton = &Singleton{}
	})
	return singleton
}
