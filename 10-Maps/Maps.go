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
