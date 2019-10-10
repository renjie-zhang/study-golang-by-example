/*
 * @Descripttion:
 * @version:
 * @Author: renjie.zhang
 * @Date: 2019-10-10 11:49:04
 * @LastEditTime: 2019-10-10 11:52:03
 */
package main

import "fmt"

func main() {
	a, b := values()
	fmt.Println("first:", a)
	fmt.Println("second:", b)

	_, c := values()
	fmt.Println("_,c:", c)
}

func values() (int, int) {
	return 7, 8
}
