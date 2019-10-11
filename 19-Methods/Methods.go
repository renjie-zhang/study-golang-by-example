/*
 * @Descripttion:
 * @version:
 * @Author: renjie.zhang
 * @Date: 2019-10-11 21:11:43
 * @LastEditTime: 2019-10-11 21:20:03
 */
package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r *rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 3, height: 4}

	fmt.Println("area:", r.area())
	fmt.Println("perim", r.perim())
	rcopy := &r
	rcopy.width = 12
	fmt.Println("after change area:",r.area())
}
