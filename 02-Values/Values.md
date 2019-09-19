# Values

在go中像值类型比如int、float32等默认值是值类型，而引用类型，比如interface、channel等默认值是nil

| 类型      | 值类型 | 引用类型 | 默认值 |
| --------- | ------ | -------- | ------ |
| int       | :o:    |          | 0      |
| float32   | :o:    |          | 0      |
| byte      | :o:    |          | 0      |
| bool      | :o:    |          | false  |
| interface |        | :o:      | <nil>  |
| channel   |        | :o:      | <nil>  |
|           |        |          |        |

```go
package main

import "fmt"

func main() {
	fmt.Println("go" + "lang")
	fmt.Println("1+1=", 1+1)
	fmt.Println("7.0/3.0=", 7.0/3.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
	var a interface{}
	var b int
	var c float32
	var d float64
	var e map[int]int
	var g bool
	var t byte
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	var f chan (int)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(t)
}

/*
输出
golang
1+1= 2
7.0/3.0= 2.3333333333333335
false
true
false
<nil>
0
0
0
map[]
<nil>
false
0
*/
```

