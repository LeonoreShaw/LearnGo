package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func responseSize(url string) {
	fmt.Println("Step1: ", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Step2: ", url)
	defer response.Body.Close()

	fmt.Println("Step3: ", url)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Step4:***$$$$$**** ", len(body), url)
}

func main() {
	go responseSize("https://baidu.com")
	go responseSize("https://jd.com")
	go responseSize("https://google.com")
	go responseSize("https://youtube.com")
	go responseSize("https://qq.com")
	time.Sleep(5 * time.Second)
}
