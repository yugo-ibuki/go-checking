package main

import (
	"fmt"
	"reflect"
)

// タグとして定義する事ができる

type Organization struct {
	Name     string `label:"名前" json:"name" validate:"required"`
	Language string `label:"言語" json:"language" validate:"required"`
}

// フォーム情報やJSONのリクエストを解析して構造体にマッピングし、問題なければバリデーションで入力値を検証し、最後にデータベースに書き込むということを1つの構造体で実現できる。

func main() {
	organization := Organization{
		Name:     "gumi",
		Language: "Elixir",
	}
	// タグを取り扱うにはreflectパッケージを利用する。
	tag := reflect.TypeOf(organization).Field(0).Tag
	fmt.Println(tag.Get("label"))      // -> 名前
	fmt.Println(tag.Get("json"))       // -> name
	fmt.Println(tag.Get("validate"))   // -> required
	fmt.Println(organization.Name)     // -> gumi
	fmt.Println(organization.Language) // -> Elixir
}
