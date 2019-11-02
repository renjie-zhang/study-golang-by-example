/*
 * @Descripttion:
 * @version:
 * @Author: renjie.zhang
 * @Date: 2019-11-02 21:43:29
 * @LastEditTime: 2019-11-02 21:47:40
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println
	t := time.Now()
	p(t.Format(time.RFC3339))

	t1, _ := time.Parse(
		time.RFC3339,
		"2019-11-01T22:08:41+00:00")
	p(t1)

	p(t.Format("3:04PM"))
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p(t.Format("2006-01-02T15:04:05.999999-07:00"))
}

/*
2019-11-02T21:47:30+08:00
2019-11-01 22:08:41 +0000 +0000
9:47PM
Sat Nov  2 21:47:30 2019
2019-11-02T21:47:30.720135+08:00
*/
