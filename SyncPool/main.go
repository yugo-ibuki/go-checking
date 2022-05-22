package main

import (
	"fmt"
	"sync"
)

// 巨大な構造体

type BigStruct struct {
	Member string
}

func main() {
	// Poolは初期化関数をNewフィールドに設定して作成する
	var pool = &sync.Pool{
		New: func() interface{} {
			return &BigStruct{}
		},
	}
	fmt.Println(pool)

	// インスタンスはGet()メソッドで取得
	// 在庫があればそれを、なければNew()を呼び出す
	b := pool.Get().(*BigStruct)
	pool.Put(b) //使い終わったらPut()でPoolに戻して次回に備える
}
