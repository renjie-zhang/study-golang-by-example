/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
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
