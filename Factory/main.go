package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
}

// 関数にNewを最初煮付けた名前にする事でファクトリー関数にする事ができる。
// 生成と同時に処理を行いたい場合、constructorとして使える。

func NewPerson(first, last string) *Person {
	return &Person{
		FirstName: first,
		LastName:  last,
	}
}

func main() {
	//person := Person{"a", "b"} // 通常
	person := NewPerson("a", "b") // ファクトリー関数で生成
	fmt.Println(person.FirstName) // a
	fmt.Println(person.LastName)  // b
}
