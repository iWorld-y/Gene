package Gene

import (
	"testing"
)

func Test_parsePattern(t *testing.T) {
	r := router{}
	url := "/123/321/aaa/bbb"
	if ans := r.parsePattern(url); !check(ans, []string{"123", "321", "aaa", "bbb"}) {
		t.Fatal("error: ", url, "got: ", ans)
	}
	url = "/123/321/*/bbb"
	if ans := r.parsePattern(url); !check(ans, []string{"123", "321", "*"}) {
		t.Fatal("error: ", url, "got: ", ans)
	}
	url = "/123/321/*name/*"
	if ans := r.parsePattern(url); !check(ans, []string{"123", "321", "*name"}) {
		t.Fatal("error: ", url, "got: ", ans)
	}
	url = "/123/321/:lang/Go"
	if ans := r.parsePattern(url); !check(ans, []string{"123", "321", ":lang", "Go"}) {
		t.Fatal("error: ", url, "got: ", ans)
	}
}

func check(parts, std []string) bool {
	if len(parts) != len(std) {
		return false
	}
	for i := 0; i < len(parts); i++ {
		if parts[i] != std[i] {
			return false
		}
	}
	return true
}

func newTestRouter() *router {
	r := newRouter()
	r.addRoute("GET", "/", nil)
	r.addRoute("GET", "/hello/:name", nil)
	r.addRoute("GET", "/learn/golang", nil)
	r.addRoute("GET", "/learn/python", nil)
	r.addRoute("GET", "/learn/java", nil)
	r.addRoute("GET", "/hi/:name", nil)
	r.addRoute("GET", "/assets/*filepath", nil)
	return r
}

func Test_getRouter(t *testing.T) {
	r := newTestRouter()
	n, params := r.getRouter("GET", "/hello/iWorld")
	if n == nil {
		t.Fatal("n = nil")
	}
	if n.pattern != "/hello/:name" {
		t.Fatal("match failed, got: ", n.pattern)
	}
	if params["name"] != "iWorld" {
		t.Fatal("params Error")
	}
	t.Logf("matched path: %s, params['name']: %s\n", n.pattern, params["name"])
}
