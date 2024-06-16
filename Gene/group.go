package Gene

import (
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc
	parent      *RouterGroup
	engine      *Engine // 所有组使用一个 engine 实例
}

func (g *RouterGroup) Group(prefix string) *RouterGroup {
	engine := g.engine
	newGroup := &RouterGroup{
		prefix: strings.Join([]string{g.prefix, prefix}, ""),
		parent: g,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (g *RouterGroup) addRoute(method, comp string, handler HandlerFunc) {
	pattern := strings.Join([]string{g.prefix, comp}, "")
	log.Printf("Route %4s - %s", method, pattern)
	g.engine.router.addRoute(method, pattern, handler)
}

func (g *RouterGroup) GET(pattern string, handler HandlerFunc) {
	g.addRoute("GET", pattern, handler)
}

func (g *RouterGroup) POST(pattern string, handler HandlerFunc) {
	g.addRoute("POST", pattern, handler)
}

// Use 添加中间件到组中
func (g *RouterGroup) Use(middlewares ...HandlerFunc) {
	g.middlewares = append(g.middlewares, middlewares...)
}

// 创建静态文件服务器
func (g *RouterGroup) createStaticHandler(relativePath string, fs http.FileSystem) HandlerFunc {
	absolutePath := path.Join(g.prefix, relativePath)
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))
	return func(ctx *Context) {
		file := ctx.Param("filepath")
		if _, err := os.Open(file); err != nil {
			ctx.Status(http.StatusNotFound)
			return
		}
		fileServer.ServeHTTP(ctx.Writer, ctx.Req)
	}
}

// Static 将本地磁盘目录映射到路由上
func (g *RouterGroup) Static(relativePath string, root string) {
	handler := g.createStaticHandler(relativePath, http.Dir(root))
	urlPattern := path.Join(relativePath, "/*filepath")
	g.GET(urlPattern, handler)
}
