package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	Url := "https://root:123456@www.baidu.com:8080/login?name=xiaoming&name=xiaoqing&age=24&age1=23#fffffff"
	u, err := url.Parse(Url)
	//  Parse函数解析Url为一个URL结构体，Url可以是绝对地址，也可以是相对地址
	//	type URL struct {
	//    Scheme   string
	//    Opaque   string    // 编码后的不透明数据
	//    User     *Userinfo // 用户名和密码信息
	//    Host     string    // host或host:port
	//    Path     string
	//    RawQuery string // 编码后的查询字符串，没有'?'
	//    Fragment string // 引用的片段（文档位置），没有'#'
	// }
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u.Scheme) // 得到Scheme
	fmt.Println(u.Path)   // 得到path  /login
	fmt.Println(u.Host)   // 得到Host路径 www.baidu.com:0000
	fmt.Println(u.Port)   // 0xd8dbc0函数地址
	fmt.Println(u.Port()) // 8080
	fmt.Println(u.RawPath)
	fmt.Println(u.String())
	fmt.Println(u.IsAbs()) // url是否是绝对距离
	user := u.User
	fmt.Println(user.Username())   // username参数
	password, b := user.Password() // Password参数 b代表是否有Password参数
	fmt.Println(password, b)

	fmt.Println(u.RawQuery)

	fmt.Println(u.Query()) // 查询query的map

	v := url.Values{}
	// 添加一个k-v值
	v.Set("page", "1")
	v.Add("name", "xiaoqing")
	v.Set("offset", "10")
	fmt.Println(v)             // output: map[name:[xiaoqing] offset:[10] page:[1]]
	fmt.Println(v.Get("name")) // output: xiaoqing
	v.Del("name")
	fmt.Println(v) // output: map[offset:[10] page:[1]]
	// 把map编码成name=xiaoming&name=xiaoqing&age=24&age1=23的形式
	v.Set("name", "xiaoming")
	v.Add("name", "xiaoqing")
	fmt.Println(v.Encode()) // output: Age=23&name=xiaoming&name=xiaoqing

}
