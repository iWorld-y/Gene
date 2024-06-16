package Gene

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

// HandlerFunc handler 方法
type HandlerFunc func(ctx *Context)

type Engine struct {
	*RouterGroup
	router   *router
	groups   []*RouterGroup     // 保存所有的组
	template *template.Template // 模板
	funcMap  template.FuncMap   // 模板渲染方法
}

func NewEngine() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

func (e *Engine) GET(pattern string, handler HandlerFunc) {
	e.router.addRoute("GET", pattern, handler)
}

func (e *Engine) POST(pattern string, handler HandlerFunc) {
	e.router.addRoute("POST", pattern, handler)
}

func (e *Engine) Run(addr string) error {
	log.Printf("Run -> %s\n", addr)
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var middlewares []HandlerFunc
	for _, group := range e.groups {
		// 判断该请求适用于哪些中间件
		if strings.HasPrefix(r.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	c := newContext(w, r)
	c.handlers = middlewares
	c.engine = e
	e.router.handler(c)
	log.Printf("ServeHTTP -> cxt: %+v", c)
	log.Printf("ServeHTTP -> Engine: %+v", e)
}

func (e *Engine) SetFuncMap(funcMap template.FuncMap) {
	e.funcMap = funcMap
}

func (e *Engine) LoadHTMLGlob(pattern string) {
	e.template = template.Must(template.New("").Funcs(e.funcMap).ParseGlob(pattern))
}
