package Gene

import (
	"log"
	"strings"
)

type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc
	parent      *RouterGroup
	engine      *Engine // 所有组使用一个 engine 实例
}

func (r *RouterGroup) Group(prefix string) *RouterGroup {
	engine := r.engine
	newGroup := &RouterGroup{
		prefix: strings.Join([]string{r.prefix, prefix}, ""),
		parent: r,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (r *RouterGroup) addRoute(method, comp string, handler HandlerFunc) {
	pattern := strings.Join([]string{r.prefix, comp}, "")
	log.Printf("Route %4s - %s", method, pattern)
	r.engine.router.addRoute(method, pattern, handler)
}
func (r *RouterGroup) GET(pattern string, handler HandlerFunc) {
	r.addRoute("GET", pattern, handler)
}
func (r *RouterGroup) POST(pattern string, handler HandlerFunc) {
	r.addRoute("POST", pattern, handler)
}
