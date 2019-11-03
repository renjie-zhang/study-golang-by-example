/*
 * @Descripttion:
 * @version:
 * @Author: renjie.zhang
 * @Date: 2019-11-03 13:02:43
 * @LastEditTime: 2019-11-03 13:04:55
 */
package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	fmt.Println("response status", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

/*
response status 200 OK
<!DOCTYPE html>
<!--STATUS OK-->
*/
