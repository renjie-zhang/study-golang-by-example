package main

import (
	"fmt"
	"time"
)

func work(id int,jobs <-chan int,result chan<- int){
	for j := range jobs{
		fmt.Println("worker",id,"processing job",j)
		time.Sleep(time.Second)
		result <- j*2
	}
}

func main(){
	jobs := make(chan int,100)
	result := make(chan int,100)

	for i:=1;i<=3;i++{
		go work(i,jobs,result)
	}

	for j :=1;j<=9;j++{
		jobs <-j
	}
	close(jobs)

	for a := 1;a<=9;a++{
		<-result
	}
}