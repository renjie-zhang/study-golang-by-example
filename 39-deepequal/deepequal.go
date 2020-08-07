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
	"reflect"
)

func main()  {
	a := map[int]string{1:"one",2:"two",3:"three"}
	b := map[int]string{1:"one",2:"two",3:"three"}
	fmt.Println("map a is equal map b: ",reflect.DeepEqual(a,b))
	c := []int{1,2,3}
	d := []int{1,2,3}
	fmt.Println("slice c is equal slice d: ",reflect.DeepEqual(c,d))
	e := []int{3,2,1}
	fmt.Println("slice d is equal slice e: ",reflect.DeepEqual(c,e))


}
