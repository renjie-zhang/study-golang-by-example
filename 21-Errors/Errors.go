/*
 * @Descripttion:
 * @version:
 * @Author: renjie.zhang
 * @Date: 2019-10-12 22:53:13
 * @LastEditTime: 2019-10-12 23:02:08
 */
package main

import (
	"fmt"
	"errors"
)

func f1(arg int) (int,error){
	if arg == 42{
		return -1,errors.New("can't work")
	}
	return arg+3,nil
}

type argError struct{
	arg int
	prob string
}

func (e *argError) Error() string{
	return fmt.Sprintf("%d-%s",e.arg,e.prob)
}

func f2(arg int) (int,error){
	if arg == 42{
		return -1,&argError{arg,"can't work"}
	}
	return arg+3,nil
}

func main() {
	for _, i := range []int{7, 42} {
        if r, e := f1(i); e != nil {
            fmt.Println("f1 failed:", e)
        } else {
            fmt.Println("f1 worked:", r)
        }
    }
    for _, i := range []int{7, 42} {
        if r, e := f2(i); e != nil {
            fmt.Println("f2 failed:", e)
        } else {
            fmt.Println("f2 worked:", r)
        }
	}
	_, e := f2(42)
    if ae, ok := e.(*argError); ok {
        fmt.Println(ae.arg)
        fmt.Println(ae.prob)
    }
}
