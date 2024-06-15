package Gene

import (
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

// addRoute 添加路由
func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	parts := r.parsePattern(pattern)
	key := strings.Join([]string{method, pattern}, "-")
	var (
		ok bool
	)
	// 若方法不存在则新建
	if _, ok = r.roots[method]; !ok {
		r.roots[method] = &node{}
	}
	// 插入路由
	r.roots[method].insert(pattern, parts, 0)
	// 注册 Handler
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

// getRouter 获取路由
func (r *router) getRouter(method string, pattern string) (*node, map[string]string) {
	root, ok := r.roots[method]
	// 若方法不存在则获取路由失败
	if !ok {
		return nil, nil
	}
	searchParts := r.parsePattern(pattern)
	params := make(map[string]string)
	n := root.search(searchParts, 0)
	if n == nil {
		return nil, nil
	}
	parts := r.parsePattern(n.pattern)
	for idx, part := range parts {
		// ":" 参数解析
		if strings.HasPrefix(part, ":") {
			params[part[1:]] = searchParts[idx]
		} else if strings.HasPrefix(part, "*") && len(part) > 1 {
			// "*" 参数解析
			// 星号之后都是参数
			params[part[1:]] = strings.Join(searchParts[idx:], "/")
			break
		}
	}
	return n, params
}
