/*
 * @Descripttion:
 * @version:
 * @Author: renjie.zhang
 * @Date: 2019-11-03 09:47:43
 * @LastEditTime: 2019-11-03 10:01:34
 */
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"bufio"
	"io"
)

func main() {
	data, err := ioutil.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	file, err := os.Open("data.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}

	b := make([]byte, 5)
	n, err := file.Read(b)
	if err != nil{
		panic(err)
	}
	fmt.Println("%d bytes: %s\n", n, string(b[:n]))

	o2, err := file.Seek(6, 0)
    if err != nil{
		panic(err)
	}
    b2 := make([]byte, 2)
    n2, err := file.Read(b2)
    if err != nil{
		panic(err)
	}
    fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))
	
	o3, err := file.Seek(6, 0)
    if err != nil{
		panic(err)
	}
    b3 := make([]byte, 2)
    n3, err := io.ReadAtLeast(file, b3, 2)
    if err != nil{
		panic(err)
	}
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))
	
	_, err = file.Seek(0, 0)
	if err != nil{
		panic(err)
	}
	
	r4 := bufio.NewReader(file)
    b4, err := r4.Peek(5)
    if err != nil{
		panic(err)
	}
    fmt.Printf("5 bytes: %s\n", string(b4))

}

/*
hello
golang
%d bytes: %s
 5 hello
2 bytes @ 6:
g
2 bytes @ 6:
g
5 bytes: hello
*/
