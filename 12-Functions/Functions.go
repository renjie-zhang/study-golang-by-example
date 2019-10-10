/*
 * @Descripttion:
 * @version:
 * @Author: renjie.zhang
 * @Date: 2019-10-10 11:45:14
 * @LastEditTime: 2019-10-10 11:47:16
 */
package main

import "fmt"

func main() {
	result := plus(3, 4)
	fmt.Println("result=", result)
}

func plus(a int, b int) int {
	return a + b
}
