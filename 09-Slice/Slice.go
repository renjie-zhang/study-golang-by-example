/*
 * @Descripttion: example of slice
 * @version:  v1
 * @Author: renjie.zhang
 * @Date: 2019-09-30 22:40:37
 * @LastEditTime: 2019-09-30 22:58:50
 */
package main

import "fmt"

func main() {
	s := make([]string, 5)

	fmt.Println("slice :", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	s[3] = "d"
	s[4] = "e"

	fmt.Println("update slice:", s)

	s = append(s, "f", "g")

	fmt.Println("after append", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("c copy from s:", c)

	a1 := s[2:5]
	a2 := s[:5]
	a3 := s[2:]

	fmt.Println(a1, a2, a3)

	twoD := make([][]int, 3)

	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)

}
