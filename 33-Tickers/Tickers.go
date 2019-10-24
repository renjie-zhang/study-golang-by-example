package main

import (
	"fmt"
	"time"
)

func main(){
	ticker := time.NewTicker(time.Millisecond * 500)

	go func(){
		for i := range ticker.C{
			fmt.Println("Tick at ",i)
		}
	}()

	time.Sleep(time.Millisecond*1600)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}