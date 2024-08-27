package gene

import (
	"fmt"
	"net/http"
	"strings"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)
type Engine struct {
	router map[string]HandlerFunc
}

func NewEngine() *Engine {
	return &Engine{
		router: make(map[string]HandlerFunc),
	}
}

// 指明这个屌毛路由使用这个处理函数
func (engine *Engine) addRouter(method string, pattern string, handler HandlerFunc) {
	key := strings.Join([]string{method, "-", pattern}, "")
	engine.router[key] = handler
}

// GET 添加一个 GET 请求
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRouter("GET", pattern, handler)
}

// POST 添加一个 POST 请求
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRouter("POST", pattern, handler)
}

// Gene, 启动!
func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path: %q\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q]: %q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}
