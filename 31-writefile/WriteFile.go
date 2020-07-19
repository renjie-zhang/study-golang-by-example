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
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	content := []byte("hello\ngo\n")
	err := ioutil.WriteFile("data.txt", content, 0644)
	if err != nil {
		panic(err)
	}

	file, err := os.Create("data2.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := file.Write(d2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := file.WriteString("writes\n")
	fmt.Printf("wrote %d bytes\n", n3)

	file.Sync()

	w := bufio.NewWriter(file)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush()
}
