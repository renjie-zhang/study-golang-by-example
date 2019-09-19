/*
 * @Descripttion:
 * @version:
 * @Author: renjie.zhang
 * @Date: 2019-09-17 09:14:34
 * @LastEditTime: 2019-09-19 16:58:46
 */
package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println("i=", i)
		i++
	}
	for j := 7; j < 9; j++ {
		fmt.Println("j=", j)
	}
	//相当于其他语言的while(true)
	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println("n=", n)
	}
	temp := [4]int{12, 4, 12, 5}
	for k, v := range temp {
		fmt.Println(k, "->", v)
	}
	//踩坑
	slice := []int{1, 2, 3}
	m := make(map[int]*int)
	for k, v := range slice {
		temp := v
		m[k] = &temp
	}
	for k, v := range m {
		fmt.Println(k, "->", *v)
	}
}

/*
1 -> 0xc0000561b8
2 -> 0xc0000561b8
0 -> 0xc0000561b8

0 -> 3
1 -> 3
2 -> 3

i= 1
i= 2
i= 3

j= 7
j= 8

loop

n= 1
n= 3
n= 5

0 -> 12
1 -> 4
2 -> 12
3 -> 5

0 -> 1
1 -> 2
2 -> 3
*/
