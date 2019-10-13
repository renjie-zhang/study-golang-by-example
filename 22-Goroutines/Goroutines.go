/*
 * @Descripttion:
 * @version:
 * @Author: renjie.zhang
 * @Date: 2019-10-13 22:24:23
 * @LastEditTime: 2019-10-13 22:27:56
 */
package main

import (
	"fmt"
	"time"
)

func f1(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f1("direct")

	go f1("goroutines")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)

	fmt.Println("done")
}
