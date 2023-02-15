package main

import (
	"fmt"
	"time"
)

func showMsg(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("msg: %v\n", msg)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	go showMsg("GoLang") // go关键字，启动协程
	go showMsg("Java")
	time.Sleep(time.Millisecond * 300)
	fmt.Println("End...")
}
