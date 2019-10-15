/*
 * @Descripttion:
 * @version:
 * @Author: renjie.zhang
 * @Date: 2019-10-15 21:15:02
 * @LastEditTime: 2019-10-15 23:05:05
 */
package main

import "fmt"

func main() {
	messages := make(chan string, 2)
	messages <- "first"
	messages <- "second"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
