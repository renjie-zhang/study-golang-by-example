/*
 * @Descripttion:
 * @version:
 * @Author: renjie.zhang
 * @Date: 2019-11-03 10:04:21
 * @LastEditTime: 2019-11-03 10:23:49
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
	err := ioutil.WriteFile("data.txt",content,0644)
	if err != nil{
		panic(err)
	}

	file,err := os.Create("data2.txt")
	if err != nil{
		panic(err)
	}

	defer file.Close()

	d2 := []byte{115, 111, 109, 101, 10}
    n2, err := file.Write(d2)
    if err != nil{
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
