package main

import (
	"container/list"
	"fmt"
)

func main() {
	// container/listを使用
	l := list.New()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	// イテレーターで表示する
	// Nextで返り値がnilでない間はループ
	for ele := l.Front(); ele != nil; ele = ele.Next() {
		fmt.Println(ele.Value)
	}
}
