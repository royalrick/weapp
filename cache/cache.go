package cache

import "time"

type Cache interface {
	// 存储数据
	Set(key string, val interface{}, timeout time.Duration)
	// 获取数据
	// 返回数据和数据是否存在
	Get(key string) (interface{}, bool)
}
