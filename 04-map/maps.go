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

import (
	"fmt"
)

func main() {
	m := make(map[string]int,10)
	fmt.Println("after new a map:", m)
	// add
	m["key1"] = 1
	m["key2"] = 2
	m["key3"] = 3
	fmt.Println("after add element to map:", m)
	// delete
	delete(m, "key2")
	fmt.Println("after delete element from map:", m)
	// update
	m["key1"] = 10
	fmt.Println("after update element from map:", m)
	// get
	fmt.Println("after get value from map: ", m["key1"])

	for key,value := range m{
		// 由于map与slice的element是共享的，
		//所以在循环中修改element会体现到原有的容器中
		fmt.Println("get a key from map： ",key)
		fmt.Println("get a value from map: ",value)
	}

	// use map implement set map[type] bool
	SetMap := make(map[int]bool,1)
	SetMap[1] = true
	SetMap[2] = true
	// select a element
	checkValueIsExisting(SetMap,2)

	// delete
	delete(SetMap,2)
	fmt.Printf("after delete %d form map\n",2)
	// select a element
	checkValueIsExisting(SetMap,2)

}

func checkValueIsExisting(m map[int]bool,n int){
	if m[n] {
		fmt.Printf("element value = %d is existing\n",n)
	}else {
		fmt.Printf("element value = %d is not existing\n",n)
	}
}

