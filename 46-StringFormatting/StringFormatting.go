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
	"os"
)

type Point struct {
	x, y int
}

func main() {
	p := Point{1, 2}
	//{1 2}
	fmt.Printf("%v\n", p)
	//{x:1 y:2}
	fmt.Printf("%+v\n", p)
	//main.Point{x:1, y:2}
	fmt.Printf("%#v\n", p)
	//main.Point
	fmt.Printf("%T\n", p)
	//true
	fmt.Printf("%t\n", true)
	//123
	fmt.Printf("%d\n", 123)
	//1110
	fmt.Printf("%b\n", 14)
	//!
	fmt.Printf("%c\n", 33)
	//1c8
	fmt.Printf("%x\n", 456)
	//78.900000
	fmt.Printf("%f\n", 78.9)
	//1.234000e+08
	//1.234000E+08
	fmt.Printf("%e\n", 123400000.0)
	fmt.Printf("%E\n", 123400000.0)
	//"string"
	fmt.Printf("%s\n", "\"string\"")
	//"\"string\""
	fmt.Printf("%q\n", "\"string\"")
	//6865782074686973
	fmt.Printf("%x\n", "hex this")
	//0xc000010090
	fmt.Printf("%p\n", &p)
	//|    12|   345|
	//|  1.20|  3.45|
	//|1.20  |3.45  |
	//|   foo|     b|
	//|foo   |b     |
	fmt.Printf("|%6d|%6d|\n", 12, 345)

	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)

	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

	fmt.Printf("|%6s|%6s|\n", "foo", "b")

	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

	s := fmt.Sprintf("a %s", "string")
	//a string
	fmt.Println(s)

	//an error
	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
