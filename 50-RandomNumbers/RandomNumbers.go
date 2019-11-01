/*
 * @Descripttion:
 * @version:
 * @Author: renjie.zhang
 * @Date: 2019-11-01 22:12:18
 * @LastEditTime: 2019-11-01 22:14:08
 */
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println()

	fmt.Println(rand.Float64())

	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()
}

/*
81,87
0.6645600532184904
7.1885709359349015,7.123187485356329
*/
