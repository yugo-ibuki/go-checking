package main

import (
	"fmt"
	"time"
)

func main() {
	tzPrint()
	sleepPrint()
}

func tzPrint() {
	now := time.Now()
	tz, _ := time.LoadLocation("America/Los_Angeles")
	future := time.Date(2015, time.October, 21, 8, 28, 0, 0, tz)
	fmt.Println(now.String())
	fmt.Println(future.Format(time.RFC3339Nano))
}

func sleepPrint() {
	// 3秒停止
	fmt.Println("3秒スリープスタート")
	time.Sleep(3 * time.Second)
	fmt.Println("3秒スリープ完了")

	// 10秒間待つ
	fmt.Println("10秒停止スタート")
	timer := time.NewTimer(10 * time.Second)
	defer timer.Stop()
	<-timer.C
	fmt.Println("10秒停止完了")
}
