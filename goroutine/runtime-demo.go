package main

import (
	"fmt"
	"runtime"
)

func main(){
	fmt.Println(runtime.NumCPU())
	runtime.GOMAXPROCS(1)
	go func(s string) {
		for i := 0; i < 2; i++ {
			fmt.Println(s)
		}
	}("world")
	for i := 0; i < 2; i++ {
		runtime.Gosched() //让出当前cpu给其他等待中的goroutine
		fmt.Println("hello")
	}
}