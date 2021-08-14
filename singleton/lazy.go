package singleton

import "sync"

var (
	lazySingleton *Singleton
	once          = &sync.Once{}
)

// GetLazyInstance 懒汉式
func GetLazyInstance() *Singleton {
	if lazySingleton == nil { // 因为这里的判断开销比 once.Do中的 atomic.LoadUint32(&o.done) == 0 更低
		once.Do(func() {
			lazySingleton = &Singleton{}
		})
	}
	return lazySingleton
}
