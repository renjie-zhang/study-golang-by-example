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
	result := plus(3, 4)
	fmt.Println("plus=", result)

	x,y := multipleValuesFunction(1,2)
	fmt.Println("multipleValuesFunction get value x：",x)
	fmt.Println("multipleValuesFunction get value y：",y)

	value := variadicFunction(1,2,3,4,5)
	fmt.Println("variadicFunction get value :",value)

	closure := closureFunction()
	fmt.Println("closure first count : ",closure())
	fmt.Println("closure second count : ",closure())
	newClosure := closureFunction()
	fmt.Println("new closure fist count:",newClosure())
}

func plus(a int, b int) int {
	return a + b
}

func multipleValuesFunction(x int ,y int) (int ,int) {
	return x, y
}

func variadicFunction(nums ...int) int {
	var result int
	for i := 0; i < len(nums); i++ {
		result = result + nums[i]
	}
	return result
}

func closureFunction() func() int{
	i := 0
	return func() int {
		i++
		return i
	}
}
