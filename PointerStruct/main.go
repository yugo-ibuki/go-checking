package main

import "fmt"

type NoCopyStruct struct {
	self  *NoCopyStruct
	Value *string
}

func NewNoCopyStruct(value string) *NoCopyStruct {
	r := &NoCopyStruct{
		Value: &value,
	}
	r.self = r
	return r
}

func (n *NoCopyStruct) String() string {
	if n != n.self {
		panic("should not copy NoCopyStruct instance without Copy() method")
	}
	return *n.Value
}

func main() {
	copy := NewNoCopyStruct("aaa")
	fmt.Println(copy)
}
