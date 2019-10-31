/*
 * @Descripttion:
 * @version:
 * @Author: renjie.zhang
 * @Date: 2019-10-31 21:31:18
 * @LastEditTime: 2019-10-31 22:23:25
 */
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type resp1 struct {
	Page   int
	Fruits []string
}

type resp2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	b, _ := json.Marshal(true)
	fmt.Println(string(b))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(3.14)
	fmt.Println(string(fltB))

	str1 := []string{"apple", "peach", "pear"}
	str2,_ := json.Marshal(str1)
	fmt.Println(string(str2))

	mapD := map[string]int{"apple": 5, "pear": 8}
	mapC,_ := json.Marshal(mapD)
	fmt.Println(string(mapC))

	result1 := &resp1{
		Page:1,
		Fruits:[]string{"apple","pear"}}
	result2,_ := json.Marshal(result1)
	fmt.Println(string(result2))

	result3 := &resp2{
        Page:1,
        Fruits:[]string{"apple","pear"}}
	result4, _ := json.Marshal(result3)
    fmt.Println(string(result4))

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	var data map[string]interface{}

	if err := json.Unmarshal(byt,&data);err !=nil{
		panic(err)
	}

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
    res := resp2{}
    json.Unmarshal([]byte(str), &res)
    fmt.Println(res)
	fmt.Println(res.Fruits[0])
	
	enc := json.NewEncoder(os.Stdout)
    d := map[string]int{"apple": 5, "lettuce": 7}
    enc.Encode(d)
}

/*
true
1
3.14
["apple","peach","pear"]
{"apple":5,"pear":8}
{"Page":1,"Fruits":["apple","pear"]}
{"page":1,"fruits":["apple","pear"]}
{1 [apple peach]}
apple
{"apple":5,"lettuce":7}
*/
