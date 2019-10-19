/*
 * @Descripttion:
 * @version:
 * @Author: renjie.zhang
 * @Date: 2019-10-19 21:42:52
 * @LastEditTime: 2019-10-19 22:01:03
 */
package main

import "fmt"

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "pass message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
