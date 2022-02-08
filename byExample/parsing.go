package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	p := fmt.Println
	u, err := url.Parse(s)

	if err != nil {
		panic(err)
	}

	p(u.Scheme)

	p(u.User)
	p(u.User.Username())
	p(u.User.Password())

	p(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	p(host, port)

	p(u.Path)
	p(u.Fragment)

	p(u.RawFragment)
	p(u.RawQuery)

	m, _ := url.ParseQuery(u.RawQuery)
	p(m)
	p(m["k"][0])
}
