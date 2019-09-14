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
}

/*
输出
i= 1
i= 2
i= 3
j= 7
j= 8
loop
n= 1
n= 3
n= 5
*/
