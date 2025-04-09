package msync

import (
	"sync"
)

// 定义一个泛型的 AnyMap，支持任意类型的键和值
type AnyMap[K comparable, V any] struct {
	m sync.Map
}

// 设置键值对
func (sm *AnyMap[K, V]) Store(key K, value V) {
	sm.m.Store(key, value)
}

// 获取键对应的值
func (sm *AnyMap[K, V]) Load(key K) (V, bool) {
	val, ok := sm.m.Load(key)
	if ok {
		return val.(V), true
	}
	var zeroValue V
	return zeroValue, false
}

// 删除键值对
func (sm *AnyMap[K, V]) Delete(key K) {
	sm.m.Delete(key)
}

// 遍历所有键值对
func (sm *AnyMap[K, V]) Range(f func(key K, value V) bool) {
	sm.m.Range(func(key, value any) bool {
		return f(key.(K), value.(V))
	})
}

func (sm *AnyMap[K, V]) LoadOrStore(key K, value V) (V, bool) {
	val, loaded := sm.m.LoadOrStore(key, value)
	return val.(V), loaded
}

func (sm *AnyMap[K, V]) LoadAndDelete(key K) (V, bool) {
	val, loaded := sm.m.LoadAndDelete(key)
	if loaded {
		return val.(V), true
	}
	var zeroValue V
	return zeroValue, false
}

func (sm *AnyMap[K, V]) Len() int {
	count := 0
	sm.m.Range(func(_, _ any) bool {
		count++
		return true
	})
	return count
}
