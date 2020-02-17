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

type student struct {
	name string
	age  int
}

func NewStudent(name string) *student {
	s := student{name: name}
	s.age = 18
	return &s
}

func main() {
	fmt.Println(student{"Bob", 13})
	fmt.Println(student{name: "Alice", age: 30})

	s := NewStudent("joker")
	fmt.Println(*s)

}
