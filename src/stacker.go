package main

import (
	"fmt"
	"goinaction/src/statistics"
	"time"
)

func main()  {
	statistics.Start()
	count, _ := fmt.Println("服务器启动成功")
	fmt.Println(count)
	go log("Hello World!")
	time.Sleep(time.Second)
}

func log(str string)  {
	fmt.Println(str)
}