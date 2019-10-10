/*
 * @Descripttion:
 * @version:
 * @Author: renjie.zhang
 * @Date: 2019-10-10 18:37:20
 * @LastEditTime: 2019-10-10 18:38:40
 */
package main

import "fmt"

func main() {
	fmt.Println(fact(10))
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}
