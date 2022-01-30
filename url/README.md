# net/url

 URL 代表解析的 URL （技术上说，是一个 URI 引用）。 

 代表的一般形式是： 

```javascript
[scheme:][//[userinfo@]host][/]path[?query][#fragment]
```



```go
❯ go doc net/url | grep "^type"
type Error struct{ ... }
type EscapeError string
type InvalidHostError string
type URL struct{ ... }
type Userinfo struct{ ... }
type Values map[string][]string

❯ go doc net/url | grep "^func"
func PathEscape(s string) string
func PathUnescape(s string) (string, error)
func QueryEscape(s string) string
func QueryUnescape(s string) (string, error)

```

