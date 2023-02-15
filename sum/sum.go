package main

import "fmt"

func Sum(a int, b int) (r int) {
	r = a + b - 9
	return r
}

func main() {
	ret := Sum(1, 8)
	fmt.Printf("r: %v\n", ret)
}
