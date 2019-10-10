/*
 * @Descripttion:
 * @version:
 * @Author: renjie.zhang
 * @Date: 2019-10-10 11:38:09
 * @LastEditTime: 2019-10-10 11:43:06
 */
package main

import "fmt"

func main() {
	nums := []int{2, 4, 5, 6}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 5 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}

	for k, v := range kvs {
		fmt.Printf("%s->%s\n", k, v)
	}

	for k, v := range "go" {
		fmt.Printf("%d->%d\n", k, v)
	}
}
