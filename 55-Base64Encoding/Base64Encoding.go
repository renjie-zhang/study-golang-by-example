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
	b64 "encoding/base64"
	"fmt"
)

func main() {
	data := "def1453!?$*&()'-=@~"
	strEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(strEnc)

	strDec, _ := b64.StdEncoding.DecodeString(strEnc)
	fmt.Println(string(strDec))
	fmt.Println()

	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}

/*
ZGVmMTQ1MyE/JComKCknLT1Afg==
def1453!?$*&()'-=@~

ZGVmMTQ1MyE_JComKCknLT1Afg==
def1453!?$*&()'-=@~
*/
