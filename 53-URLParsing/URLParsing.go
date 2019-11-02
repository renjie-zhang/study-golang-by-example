/*
 * @Descripttion:
 * @version:
 * @Author: renjie.zhang
 * @Date: 2019-11-02 22:46:04
 * @LastEditTime: 2019-11-02 22:48:42
 */
package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Println(u.Scheme)

	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])
}

/*
postgres
user:pass
user
pass
host.com:5432
host.com
5432
k=v
map[k:[v]]
v
*/
