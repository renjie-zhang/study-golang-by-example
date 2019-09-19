# For循环

在go语言中只存在一种循环，也就是for循环，如果有着while（true）的需求，可以使用for{}这样的形式来替换。

```go
package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println("i=", i)
		i++
	}
	for j := 7; j < 9; j++ {
		fmt.Println("j=", j)
	}
	//相当于其他语言的while(true)
	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println("n=", n)
	}
	temp := [4]int{12, 4, 12, 5}
	for k, v := range temp {
		fmt.Println(k, "->", v)
	}

}
```
## Attention






## 坑1:angry:

```go
slice := []int{1, 2, 3}
m := make(map[int]*int)
for k, v := range slice {
	m[k] = &v
}
for k, v := range m {
	fmt.Println(k, "->", *v)
}
```

猜猜会输出什么？？:happy:

```json
0 -> 3
1 -> 3
2 -> 3
```

这是为什么？:cry: 为什么不是我们想的那样。再来看看输出的地址是多少。

```go
1 -> 0xc0000561b8
2 -> 0xc0000561b8
0 -> 0xc0000561b8
```

好像明白了什么:laughing:

原来问题在循环变量的作用域。for循环语句引入了新的词法块，循环变量v这个在语法块中被声明。在该循环中的所有函数值都共享相同的循环变量。需要注意的是函数值记录的是循环变量的内存地址，而不是循环变量某一时刻的值。所以说最终的操作结束，也就是for循环结束之后，最后一个赋值到内存地址中，前面的两个元素也是指向同一个地址，所以全部为3。***此问题在Go语言圣经第5.6章匿名函数中有讲解***。

正确写法：

```go
slice := []int{1, 2, 3}
m := make(map[int]*int)
for k, v := range slice {
	temp := v
	m[k] = &temp
}
for k, v := range m {
	fmt.Println(k, "->", *v)
}
```

