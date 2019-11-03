/*
 * @Descripttion:
 * @version:
 * @Author: renjie.zhang
 * @Date: 2019-11-03 13:07:17
 * @LastEditTime: 2019-11-03 13:12:16
 */
package main

import (
	"fmt"
	"net/http"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func header(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	http.HandleFunc("/", greet)
	http.HandleFunc("/headers", header)
	http.ListenAndServe(":8080", nil)
}

/*
curl localhost:8080/hello
Hello World! 2019-11-03 13:10:52.663499 +0800 CST m=+20.377621001
*/
