/*
 * @Descripttion:
 * @version:
 * @Author: renjie.zhang
 * @Date: 2019-10-11 20:58:45
 * @LastEditTime: 2019-10-11 21:02:38
 */
package main

import "fmt"

func main() {
	i := 1

	fmt.Println("init:",i)

	zeroValue(i)
	fmt.Println("zeroValue",i)

	zeroPointer(&i)
	fmt.Println("zeroPointer",i)

	fmt.Println(&i)
}

func zeroValue(value int) {
	value = 0
}

func zeroPointer(pointer *int) {
	*pointer = 0
}
