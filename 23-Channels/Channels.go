/*
 * @Descripttion:
 * @version:
 * @Author: renjie.zhang
 * @Date: 2019-10-13 22:32:28
 * @LastEditTime: 2019-10-13 22:36:59
 */
package main

import "fmt"

func main() {
	messages := make(chan string)

	go func() {
		messages <- "ping"
	}()

	go func() {
		messages <- "pong"
	}()

	for i := 0; i < 2; i++ {
		fmt.Println("messages", <-messages)
	}
}
