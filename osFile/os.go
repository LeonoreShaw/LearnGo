package main

import (
	"fmt"
	"os"
)

// 创建文件
func creatFile() {
	f, err := os.Create("c.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("f.Name(): %v\n", f.Name())
	}
}

// 创建目录
func makeDir() {
	err := os.Mkdir("B", os.ModePerm)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	err2 := os.MkdirAll("A/b/c", os.ModePerm)
	if err2 != nil {
		fmt.Printf("err2: %v\n", err2)
	}

}

// 删除目录或者文件
func remove() {
	err := os.Remove("a.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	err2 := os.RemoveAll("A")
	if err2 != nil {
		fmt.Printf("err2: %v\n", err2)
	}

}

// 工作目录，临时目录
func wd() {
	dir, _ := os.Getwd()
	fmt.Printf("dir: %v\n", dir)
	// os.Chdir("d:/")
	// dir, _ = os.Getwd()
	// fmt.Printf("dir2: %v\n", dir)
	s := os.TempDir()
	fmt.Printf("s: %v\n", s)
}

// 重命名文件
func rename() {
	err := os.Rename("test.c", "test.java")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

// 读文件
func readFile() {
	b, _ := os.ReadFile("c.txt")
	fmt.Printf("b: %v\n", string(b[:]))
}

// 写文件（覆盖重写）
func wriiteFile() {
	os.WriteFile("c.txt", []byte("OK"), os.ModePerm)
}

func main() {
	// creatFile()
	// makeDir()
	// remove()
	// wd()
	// rename()
	// readFile()
	wriiteFile()
}
