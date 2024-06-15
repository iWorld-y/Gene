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
