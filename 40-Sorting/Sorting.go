/*
 * @Descripttion:
 * @version:
 * @Author: renjie.zhang
 * @Date: 2019-10-28 22:45:38
 * @LastEditTime: 2019-10-28 22:46:22
 */
package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"t","d","r","y"}
	sort.Strings(strs)
	fmt.Println(strs)

	ints := []int{1,4,53,56,23,45}
	sort.Ints(ints)
	fmt.Println(ints)
}
