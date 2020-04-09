package model

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidNode = errors.New("invalid node")
)

type TreeNode struct {
	Parant   *TreeNode
	Children []*TreeNode
	Value    string
}

func (n *TreeNode) InsertValues(values []string) error {
	len := len(values)
	if len == 0 {
		return nil
	}

	for _, node := range n.Children {
		if node.GetValue() == values[0] {
			if len == 1 {
				return nil
			} else {
				return node.InsertValues(values[1:])
			}
		}
	}

	node := &TreeNode{
		Value: values[0],
	}
	n.AddChild(node)
	if len > 1 {
		return node.InsertValues(values[1:])
	}
	return nil
}

func (n *TreeNode) SetValue(v string) bool {
	n.Value = v
	return true
}

func (n *TreeNode) GetValue() string {
	return n.Value
}

func (n *TreeNode) AddChild(tn *TreeNode) error {
	if tn == nil {
		return ErrInvalidNode
	}

	tn.setParent(n)
	n.Children = append(n.Children, tn)
	return nil
}

func (n *TreeNode) GetChildren() []*TreeNode {
	return n.Children
}

func (n *TreeNode) setParent(tn *TreeNode) error {
	if n == nil {
		return ErrInvalidNode
	}
	n.Parant = n
	return nil
}

func (n *TreeNode) GetParant() *TreeNode {
	return n.Parant
}

func (n *TreeNode) String(level int) string {
	var res string
	res = fmt.Sprintf("level:%d, str:%s\n", level, n.Value)
	for _, child := range n.Children {
		res += child.String(level + 1)
	}
	return res
}

type UrlTree struct {
	root *TreeNode
}

func NewUrlTree(rootValue string) *UrlTree {
	return &UrlTree{
		root: &TreeNode{
			Value: rootValue,
		},
	}
}

func (t *UrlTree) Insert(rawurl string) error {
	uh, err := ParseUrl(rawurl)
	if err != nil {
		return err
	}

	sections := uh.GetPathSection()
	return t.root.InsertValues(sections)
}

func (t *UrlTree) Delete(rawurl string) error {
	return nil
}

func (t *UrlTree) Update(origin, target string) error {
	return nil
}

func (t *UrlTree) Get(rawurl string) (*TreeNode, error) {
	return nil, nil
}

func (t *UrlTree) String() string {
	return t.root.String(0)
}
