package main

import (
	"fmt"
	"math"
)

const s string = "constants"

func main() {
	fmt.Println(s)

	const n = 40000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}

/*
输出：
constants
7.5e+15
7500000000000000
0.9465396567857337
*/
