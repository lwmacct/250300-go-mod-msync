package msync

import "sync"

// EventMap 是一个支持事件通知的键值对存储结构
type EventMap[K comparable, V any, E any] struct {
	m         sync.Map
	mu        sync.RWMutex
	callBacks []func(event E)
}

// 设置键值对
func (t *EventMap[K, V, E]) Store(key K, value V) {
	t.m.Store(key, value)
}

// 获取键对应的值
func (t *EventMap[K, V, E]) Load(key K) (V, bool) {
	val, ok := t.m.Load(key)
	if ok {
		return val.(V), true
	}
	var zeroValue V
	return zeroValue, false
}

// 删除键值对
func (t *EventMap[K, V, E]) Delete(key K) {
	t.m.Delete(key)
}

// 遍历所有键值对
func (t *EventMap[K, V, E]) Range(f func(key K, value V) bool) {
	t.m.Range(func(key, value any) bool {
		return f(key.(K), value.(V))
	})
}

func (t *EventMap[K, V, E]) LoadOrStore(key K, value V) (V, bool) {
	val, loaded := t.m.LoadOrStore(key, value)
	return val.(V), loaded
}

func (t *EventMap[K, V, E]) LoadAndDelete(key K) (V, bool) {
	val, loaded := t.m.LoadAndDelete(key)
	if loaded {
		return val.(V), true
	}
	var zeroValue V
	return zeroValue, false
}

func (t *EventMap[K, V, E]) AddCallback(callback func(event E)) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.callBacks = append(t.callBacks, callback)
}

func (t *EventMap[K, V, E]) NotifyCallbacks(event E) {
	t.mu.RLock()
	defer t.mu.RUnlock()
	for _, cb := range t.callBacks {
		cb(event)
	}
}
