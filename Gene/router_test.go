package Gene

import (
	"testing"
)

func Test_parsePattern(t *testing.T) {
	r := router{}
	url := "/123/321/aaa/bbb"
	if !check(r.parsePattern(url), []string{"123", "321", "aaa", "bbb"}) {
		t.Fatal("error: ", url)
	}
	url = "/123/321/*/bbb"
	if !check(r.parsePattern(url), []string{"123", "321", "*"}) {
		t.Fatal("error: ", url)
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
