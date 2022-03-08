package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//demo1()
	//demo2()
	//demo3()
	demo4()
}

func demo1() {
	go hello()
	fmt.Println("world")
	//主进程结束,hello未来得及打印
}
func hello() {
	fmt.Println("hello")
}

func demo2() {
	//引入channel控制主进程结束
	var ch chan bool = make(chan bool)
	go func() {
		fmt.Println("hello")
		ch <- true
	}()
	fmt.Println("world")
	<-ch
}

func demo3()  {
	var wg sync.WaitGroup
	//WaitGroup加锁
	for i := 0; i < 10; i++ {
		wg.Add(1) // 启动一个goroutine就登记+1
		go func(num int) {
			defer wg.Done()
			fmt.Println("Hello Goroutine!", num)
		}(i)
	}
	wg.Wait() // 等待所有登记的goroutine都结束
}

func demo4()  {
	// 测试主协程退出其他任务是否执行
	go func() {
		i := 0
		for {
			i++
			fmt.Printf("new goroutine: i = %d\n", i)
			time.Sleep(time.Second)
		}
	}()
	i := 0
	for {
		i++
		fmt.Printf("main goroutine: i = %d\n", i)
		time.Sleep(time.Second)
		if i == 2 {
			break
		}
	}
	//main goroutine: i = 1
	//new goroutine: i = 1
	//new goroutine: i = 2
	//main goroutine: i = 2
	//主协程结束其他任务不执行
}