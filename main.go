package main

import (
	"fmt"

	"github.com/lwmacct/250300-go-mod-msync/pkg/msync"
)

func main() {
	// 使用 SyncMap 存储字符串 -> int 类型的键值对
	sm1 := &msync.AnyMap[string, int]{}
	sm1.Store("apple", 5)
	sm1.Store("banana", 10)

	// 加载并打印值
	if value, ok := sm1.Load("apple"); ok {
		fmt.Println("apple:", value)
	}
	if value, ok := sm1.Load("banana"); ok {
		fmt.Println("banana:", value)
	}

	// 使用 SyncMap 存储 int -> string 类型的键值对
	sm2 := &msync.AnyMap[int, string]{}
	sm2.Store(1, "one")
	sm2.Store(2, "two")

	// 遍历所有键值对
	sm2.Range(func(key int, value string) bool {
		fmt.Printf("%d: %s\n", key, value)
		return true
	})
}
