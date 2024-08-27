package gene

import "net/http"

type Context struct {
	Writer http.ResponseWriter
	Req    *http.Request

	Path   string
	Method string

	StatusCode int64
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

// 返回请求中指定键 (key) 的第一个值. 如果该键不存在, 返回空字符串.
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}
