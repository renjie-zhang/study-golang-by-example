package main

import (
	"fmt"
	"time"
)

func main(){
	//速度限制，每200ms处一请求
	request := make(chan int,5)
	for i:= 1;i<=5;i++{
		request <- i
	}
	close(request)

	limiter := time.Tick(time.Millisecond*200)

	for req := range request{
		<- limiter
		fmt.Println("request",req,time.Now())
	}
	//速度限制，每次处理3个一组
	burstyLimiter := make(chan time.Time,3)

	for i := 0;i<3;i++{
		burstyLimiter <- time.Now()
	}

	go func(){
		for t := range time.Tick(time.Millisecond *200){
			burstyLimiter <- t
		}
	}()

	burstyRequest := make(chan int,5)
	for i := 1;i<=5;i++{
		burstyRequest <- i
	}
	close(burstyRequest)
	for req := range burstyRequest{
		<- burstyLimiter
		fmt.Println("request",req,time.Now())
	}

}