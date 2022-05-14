package main

import "fmt"

// 四角形の構造体とそれの面積を求めるメソッド

type Square struct {
	width  float64
	height float64
}

// Square タイプにArea関数を追加

func (s Square) Area() float64 {
	return s.width * s.height
}

func (s *Square) Reshape(w float64, h float64) {
	s.width = w
	s.height = h
}

func main() {
	square := Square{3.0, 4.0}
	fmt.Println(square)
	// output: {3 4}

	// ポインタレシーバ：ポインタとして引数で渡すので、関数内でオブジェクトの値を変更できる
	(&square).Reshape(5.0, 6.0)
	fmt.Println(square)
	// output: {5 6}
}
