/*
 * @Descripttion:
 * @version:
 * @Author: renjie.zhang
 * @Date: 2019-10-10 18:31:51
 * @LastEditTime: 2019-10-10 18:34:41
 */
package main

import "fmt"

func main() {
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())

	new := intSeq()

	fmt.Println(new())

}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
