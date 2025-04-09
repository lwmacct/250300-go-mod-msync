package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/lwmacct/250300-go-mod-msync/pkg/msync"
)

func main() {
	testEventMap()

	fmt.Println("--------------------------------")
	testAnyMap()
}

type EventData struct {
	Key   string
	Value int
}

func testEventMap() {
	eventMap := &msync.EventMap[string, int, *EventData]{}
	eventMap.Store("a", 10)
	eventMap.Store("b", 20)
	eventMap.Store("c", 30)

	eventMap.AddCallback(func(event *EventData) {
		fmt.Println("回调 1", event)
	})

	eventMap.AddCallback(func(event *EventData) {
		fmt.Println("回调 2", event)
	})

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		time.Sleep(3 * time.Second)
		eventMap.NotifyCallbacks(&EventData{Key: "事件a", Value: 100})
		wg.Done()
	}()

	fmt.Println("等待事件触发")
	wg.Wait()

	eventMap.Store("a", 100)

}

func testAnyMap() {
	// 创建一个 AnyMap 实例，键是字符串类型，值是整数类型
	stringToIntMap := &msync.AnyMap[string, int]{}

	// 使用 Store 方法存储键值对
	stringToIntMap.Store("a", 10)
	stringToIntMap.Store("b", 20)
	stringToIntMap.Store("c", 30)

	// 使用 Load 方法获取键对应的值
	value, ok := stringToIntMap.Load("b")
	if ok {
		fmt.Println("Key 'b' has value:", value)
	} else {
		fmt.Println("Key 'b' not found")
	}

	// 使用 Delete 方法删除键值对
	stringToIntMap.Delete("a")

	// 使用 Range 方法遍历所有键值对
	stringToIntMap.Range(func(key string, value int) bool {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
		return true // 返回 true 继续遍历
	})

}
