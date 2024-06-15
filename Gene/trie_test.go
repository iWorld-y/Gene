package Gene

import (
	"testing"
)

func Test_insert(t *testing.T) {
	n := node{}
	n.insert("/123/321", []string{"123", "321"}, 0)
	t.Logf("%+v", n)
}
