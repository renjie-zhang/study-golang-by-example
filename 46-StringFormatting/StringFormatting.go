/*
 * @Descripttion: 
 * @version: 
 * @Author: renjie.zhang
 * @Date: 2019-10-29 21:27:15
 * @LastEditTime: 2019-10-29 21:39:33
 */
package main

import (
	"fmt"
	"os"
)

type Point struct{
	x,y int
}
func main() {
	p := Point{1,2}
	//{1 2}
	fmt.Printf("%v\n",p)
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
