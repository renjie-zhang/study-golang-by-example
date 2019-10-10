/*
 * @Descripttion:
 * @version:
 * @Author: renjie.zhang
 * @Date: 2019-09-27 20:50:02
 * @LastEditTime: 2019-09-30 16:45:28
 */
package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set:", s)
	fmt.Println("get", s[2])

	fmt.Println("length:", len(s))

	s = append(s, "d")
	s = append(s, "e")

	fmt.Println("after append", s)
	fmt.Println("after append length", len(s))

	c := make([]string, len(s))
	copy(c, s)

	fmt.Println("s copy from c", c)

	a1 := s[:5]
	a2 := s[2:5]
	a3 := s[2:]
	fmt.Println(a1, a2, a3)
}
