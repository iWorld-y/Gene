package Gene

import (
	"testing"
)

func Test_insert(t *testing.T) {
	n := node{}
	n.insert("/123/321/aaa/bbb", []string{"123", "321", "aaa", "bbb"}, 0)
	n.insert("/123/333/aaa/bbb", []string{"123", "333", "aaa", "bbb"}, 0)
	if n.search([]string{"123", "321", "aaa", "bbb"}, 0) == nil {
		t.Fatal("n.search /123/321/aaa/bbb error")
	}
	if n.search([]string{"123", "321", "aaa"}, 0) != nil {
		t.Fatal("n.search /321/aaa/bbb error")
	}

}
