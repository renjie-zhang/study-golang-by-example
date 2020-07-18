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
	s := make([]string, 3)
	fmt.Println("default value:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("get", s)
	fmt.Println("length:", len(s))
	s = append(s, "d")
	s = append(s, "e")
	fmt.Println("after append", s)
	fmt.Println("after append length", len(s))

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("s copy from c", c)
	a1 := s[:5]
	a2 := s[2:5]
	a3 := s[2:]
	fmt.Println(a1, a2, a3)

	s[1] = "z"
	fmt.Println("update s[1] set value is z",s)
	fmt.Println(c)

}
