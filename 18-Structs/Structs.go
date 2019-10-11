/*
 * @Descripttion:
 * @version:
 * @Author: renjie.zhang
 * @Date: 2019-10-11 21:04:33
 * @LastEditTime: 2019-10-11 21:10:56
 */
package main

import "fmt"

type student struct {
	name string
	age  int
}

func NewStudent(name string) *student {
	s := student{name: name}
	s.age = 18
	return &s
}

func main() {
	fmt.Println(student{"Bob", 13})
	fmt.Println(student{name: "Alice", age: 30})

	s := NewStudent("joker")
	fmt.Println(*s)

}
