package model

import (
	"errors"
)

const (
	PreOrderTranverse   = 0
	LevelOrderTraversal = 1
)

var (
	ErrInvalidNode  = errors.New("invalid node")
	ErrParamInvalid = errors.New("param invalid")
)

//TreeNode
type TreeNode struct {
	Parant   *TreeNode
	Children []*TreeNode
	Value    string
	Level    int
}

type NodeValue struct {
	Value string
	Level int
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
		Level: n.Level + 1,
	}
	n.addChild(node)
	if len > 1 {
		return node.InsertValues(values[1:])
	}
	return nil
}

func (n *TreeNode) addChild(c *TreeNode) error {
	if c == nil {
		return ErrInvalidNode
	}

	c.Parant = n
	n.Children = append(n.Children, c)
	return nil
}

func (n *TreeNode) GetValue() string {
	return n.Value
}

func (n *TreeNode) SetValue(v string) error {
	n.Value = v
	return nil
}

func (n *TreeNode) String() string {
	return ""
}

func (n *TreeNode) Traversal(method int) ([]NodeValue, error) {
	var res []NodeValue
	switch method {
	case PreOrderTranverse:
		err := n.preOrderTraversal(&res)
		return res, err
	case LevelOrderTraversal:
		return n.levelOrderTraversal()
	default:
		return n.preOrderTraversalEx()
	}
	return n.preOrderTraversalEx()
}

func (n *TreeNode) preOrderTraversal(res *[]NodeValue) error {
	if res == nil {
		return ErrParamInvalid
	}

	*res = append(*res, NodeValue{n.Value, n.Level})
	for _, child := range n.Children {
		child.preOrderTraversal(res)
	}
	return nil
}

func (n *TreeNode) preOrderTraversalEx() ([]NodeValue, error) {
	var res []NodeValue
	s := NewStack()
	s.Push(n)
	for s.Len() > 0 {
		temp := s.Pop()
		if tn, ok := temp.(*TreeNode); ok {
			res = append(res, NodeValue{tn.Value, tn.Level})
			for _, c := range tn.Children {
				s.Push(c)
			}
		}
	}
	return res, nil
}

func (n *TreeNode) levelOrderTraversal() ([]NodeValue, error) {
	var res []NodeValue
	q := NewQueue()
	q.Push(n)
	for q.Len() > 0 {
		temp := q.Pop()
		if tn, ok := temp.(*TreeNode); ok {
			res = append(res, NodeValue{tn.Value, tn.Level})
			for _, c := range tn.Children {
				q.Push(c)
			}
		}
	}
	return res, nil
}

//UrlTree
type UrlTree struct {
	root *TreeNode
}

func NewUrlTree(rootValue string) *UrlTree {
	return &UrlTree{
		root: &TreeNode{
			Value: rootValue,
			Level: 0,
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

func (t *UrlTree) Update(originUrl, targetUrl string) error {
	return nil
}

func (t *UrlTree) Get(rawurl string) (*TreeNode, error) {
	return nil, nil
}

func (t *UrlTree) String() string {
	return t.root.String()
}

func (t *UrlTree) Traversal(method int) ([]NodeValue, error) {
	return t.root.Traversal(method)
}
