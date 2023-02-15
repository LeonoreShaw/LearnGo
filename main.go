/*
 		学习GoLang
 */
package main  					// 主包

import (
	"LearnGo/userPackage"
	"fmt"
)

var i = varFunc()				//变量初始化，先于init函数执行

func init() {					//init函数，先于main函数执行
	fmt.Println("init...")
	fmt.Println(i+2)
}

func varFunc() int {
	fmt.Println("var...")
	return 0
}

func main()  { 						// 主函数（程序入口）
	v:=userPackage.Hello()
	fmt.Printf("v: %T\n", v)
	defer fmt.Printf("end end %v\n", v)
	defer fmt.Printf("\n倒计时..........\n")
	fmt.Println("OK")

	fmt.Println("指针。。。")
	var ip *int
	fmt.Printf("ip: %v\n", ip) // nil
	fmt.Printf("ip: %T\n", ip)
	var i int = 100
	ip = &i
	fmt.Printf("ip: %v\n", ip)
	fmt.Printf("*ip: %v\n", *ip)
	fmt.Printf("*ip: %T\n", *ip)

}