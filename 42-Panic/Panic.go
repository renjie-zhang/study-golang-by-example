/*
 * @Descripttion: 
 * @version: 
 * @Author: renjie.zhang
 * @Date: 2019-10-28 22:55:22
 * @LastEditTime: 2019-10-28 22:56:29
 */
package main

import (
	"os"
)

func main(){
	panic("a problem")

	_,err := os.Create("/temp/file")
	if err != nil{
		panic(err)
	}
}