package main

import (
	"fmt"
	"time"
)


func slapper(slapper <-chan int, paer chan<- int) {//2個channel 參數,slapper(整數),paer(整數)
    for {
        <-slapper//接收參數slapper
        fmt.Println("slap")//印出slap
        time.Sleep(time.Second)//暫停1秒
        paer <- 1//傳出參數paer(為1)
    }
}

// The ponger prints a pong and waits for a ping
func paer(slapper chan<- int, paer <-chan int) {//2個channel 參數,slapper(整數),paer(整數)
    for {
        <-paer//接收參數paer
        fmt.Println("pa")//印出pa
        time.Sleep(time.Second)//暫停1秒
        slapper <- 1//傳出參數slapper(為1)
    }
}

func main() {
    slap := make(chan int)//把channel中的slap參數初始化
    pa := make(chan int)//把channel中的pa參數初始化

    go slapper(slap, pa)//做goroutine
    go paer(slap, pa)//做goroutine


    slap <- 1//給slapper中的slaph參數賦值，開始跑goroutine

    for {
        // slapper跟paer的goroutine中間進行休息
        time.Sleep(time.Second)
    }
}