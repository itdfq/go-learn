package singleton

import "sync"

/**
单例模式需要一个全局构造函数，这个构造函数会返回一个私有的对象，无论何时调用，它总是返回相同的对象。
*/

type singleton struct {
	Value int
}

type Singleton interface {
	GetValue() int
}

func (s singleton) GetValue() int {
	return s.Value
}

var (
	instance *singleton
	once     sync.Once
)

// GetInstance 构造方法，用于获取单例模式对象
func GetInstance(v int) Singleton {
	//只会调用一次
	once.Do(func() {
		instance = &singleton{Value: v}
	})
	return instance
}
