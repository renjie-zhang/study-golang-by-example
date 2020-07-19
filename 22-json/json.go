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
	"encoding/json"
	"fmt"
	"os"
)

type resp1 struct {
	Page   int
	Fruits []string
}

type resp2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	b, _ := json.Marshal(true)
	fmt.Println(string(b))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(3.14)
	fmt.Println(string(fltB))

	str1 := []string{"apple", "peach", "pear"}
	str2, _ := json.Marshal(str1)
	fmt.Println(string(str2))

	mapD := map[string]int{"apple": 5, "pear": 8}
	mapC, _ := json.Marshal(mapD)
	fmt.Println(string(mapC))

	result1 := &resp1{
		Page:   1,
		Fruits: []string{"apple", "pear"}}
	result2, _ := json.Marshal(result1)
	fmt.Println(string(result2))

	result3 := &resp2{
		Page:   1,
		Fruits: []string{"apple", "pear"}}
	result4, _ := json.Marshal(result3)
	fmt.Println(string(result4))

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	var data map[string]interface{}

	if err := json.Unmarshal(byt, &data); err != nil {
		panic(err)
	}

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := resp2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
