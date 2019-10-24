package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){
	var wg sync.WaitGroup

	for i := 1;i<=5;i++{
		wg.Add(1)
		go worker(i,&wg)
	}

	wg.Wait()
}

func worker(id int,wg *sync.WaitGroup){
	fmt.Printf("work %d starting\n",id)

	time.Sleep(time.Second)

	fmt.Printf("work %d done\n",id)

	wg.Done()
}
