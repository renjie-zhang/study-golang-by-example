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
	"github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/runtime/protoimpl"
)

//resource https://developers.google.com/protocol-buffers/docs/gotutorial
//		   https://github.com/golang/protobuf

func main() {
	user := &User{
		state:         protoimpl.MessageState{},
		sizeCache:     0,
		unknownFields: nil,
		Name:          "renjie",
		Male:          "male",
		Address:       []string{"sichuan", "chongqing"},
	}

	marshal, err := proto.Marshal(user)
	if err != nil {
		panic(err)
	}
	newUser := &User{}
	err = proto.Unmarshal(marshal, newUser)
	if err != nil {
		panic(err)
	}
	fmt.Printf("user name is：%s\n", user.Name)
	fmt.Printf("new user name is:%s\n", newUser.Name)

	/**
	  user name is：renjie
	  new user name is:renjie
	*/

}
