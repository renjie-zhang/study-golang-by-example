/*
 * @Descripttion:
 * @version:
 * @Author: renjie.zhang
 * @Date: 2019-11-02 23:02:31
 * @LastEditTime: 2019-11-02 23:06:56
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
