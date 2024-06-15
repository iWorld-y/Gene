package Gene

import (
	"strings"
)

type node struct {
	pattern  string  // 一个路由
	part     string  // 路由的一部分
	children []*node //路由的子节点
	isWild   bool    // 是否粗略匹配, 若 part 含有 ":" 或 "*" 时则为粗略匹配
}

// 第一个匹配成功的节点
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// 所有匹配成功的节点
func (n *node) matchChildren(part string) []*node {
	var nodes []*node
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// insert 插入路由
func (n *node) insert(pattern string, parts []string, height int) {
	// 当节点深度达到 parts 的长度时, 已无法再继续插入
	// 将完整的路由存放于该节点
	if len(parts) == height {
		// 在查找时, 若某节点存放着完整的 url 即路由匹配成功
		n.pattern = pattern
		return
	}
	// 取当前节点用于插入
	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		// 不存在该节点, 则新建
		child = &node{
			part:   part,
			isWild: strings.HasPrefix(part, "*") || strings.HasPrefix(part, ":"),
		}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}
