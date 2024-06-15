package Gene

import (
	"log"
	"net/http"
	"strings"
)

type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Router %4s - %s", method, pattern)
	key := strings.Join([]string{method, pattern}, "-")
	r.handlers[key] = handler
}

func (r *router) handler(c *Context) {
	key := strings.Join([]string{c.Methon, c.Path}, "-")
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}

// parsePattern 划分 pattern 为 parts, 且只允许一个 "*" 片段存在
func (r *router) parsePattern(pattern string) []string {
	split := strings.Split(pattern, "/")
	var parts []string
	for _, part := range split {
		if part != "" {
			parts = append(parts, part)
			if strings.HasPrefix(part, "*") {
				break
			}
		}
	}
	return parts
}

// addRouter 添加路由
func (r *router) addRouter(method string, pattern string, handler HandlerFunc) {
	parts := r.parsePattern(pattern)
	key := strings.Join([]string{method, pattern}, "-")
	var (
		root *node
		ok   bool
	)
	// 若方法不存在则新建
	if root, ok = r.roots[method]; !ok {
		root = &node{}
	}
	// 插入路由
	root.insert(pattern, parts, 0)
	// 注册 Handler
	r.handlers[key] = handler
}
