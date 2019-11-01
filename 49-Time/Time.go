/*
 * @Descripttion:
 * @version:
 * @Author: renjie.zhang
 * @Date: 2019-11-01 22:08:56
 * @LastEditTime: 2019-11-01 22:10:54
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	now := time.Now()
	p(now)

	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	diff := now.Sub(then)
	p(diff)
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	p(then.Add(diff))
	p(then.Add(-diff))
}

/*
2019-11-01 22:10:58.6701895 +0800 CST m=+0.005999201
2009-11-17 20:34:58.651387237 +0000 UTC
2009
November
UTC
87257h36m0.018802263s
87257.60000522285
5.235456000313371e+06
3.141273600188023e+08
314127360018802263
2019-11-01 14:10:58.6701895 +0000 UTC
*/
