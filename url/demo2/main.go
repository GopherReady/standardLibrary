package main

import (
	"fmt"
	"net/url"
)

func main() {
	u := &url.URL{
		Scheme:   "https",
		User:     url.UserPassword("admin", "admin"),
		Host:     "example.com",
		Path:     "foo/bar",
		RawQuery: "x=1&y=2",
		Fragment: "anchor",
	}
	fmt.Println(u.String())
	u.Opaque = "opaque"
	fmt.Println(u.String())

	fmt.Println(u.User.Username())
	fmt.Println(u.User.Password())
	fmt.Println(url.User("sd"))
}
