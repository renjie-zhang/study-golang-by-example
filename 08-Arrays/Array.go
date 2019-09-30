/*
 * @Descripttion:
 * @version:
 * @Author: renjie.zhang
 * @Date: 2019-09-27 20:42:24
 * @LastEditTime: 2019-09-27 20:49:44
 */
package main

import "fmt"

func main() {

	//create a array
	var temp [5]int
	//default an array is default-value
	fmt.Println("temp:", temp)

	temp[2] = 100
	fmt.Println("set", temp)
	fmt.Println("get", temp[2])
	//len() get the length of array
	fmt.Println(len(temp))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("twoD:", twoD)
}

/*
out
temp: [0 0 0 0 0]
set [0 0 100 0 0]
get 100
5
[1 2 3 4 5]
twoD: [[0 1 2] [1 2 3]]
*/
