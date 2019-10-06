/*
 * @Descripttion:
 * @version: maps
 * @Author: renjie.zhang
 * @Date: 2019-10-06 22:36:56
 * @LastEditTime: 2019-10-06 22:42:33
 */
package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["key1"] = 1
	m["key2"] = 2

	fmt.Println("map:", m)

	v1 := m["key1"]
	fmt.Println("v1 from map：", v1)

	fmt.Println("length=", len(m))

	delete(m, "key2")

	fmt.Println("after delete from m,m=", m)

}

/*
输出
map: map[key1:1 key2:2]
v1 from map： 1
length= 2
after delete from m,m= map[key1:1]
*/
