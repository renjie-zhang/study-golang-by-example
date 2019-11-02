/*
 * @Descripttion:
 * @version:
 * @Author: renjie.zhang
 * @Date: 2019-11-02 21:51:42
 * @LastEditTime: 2019-11-02 21:54:55
 */
package main

import (
	"fmt"
	"strconv"
)

func main() {
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	i,_ := strconv.ParseInt("123",0,64)
	fmt.Println(i)

	d,_ := strconv.ParseInt("0x1c8",0,64)
	fmt.Println(d)

	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)
	
	k, _ := strconv.Atoi("135")
	fmt.Println(k)
	
	_, e := strconv.Atoi("wat")
    fmt.Println(e)
}

/*
1.234
123
456
789
135
strconv.Atoi: parsing "wat": invalid syntax
*/

