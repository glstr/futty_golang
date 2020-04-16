package model

import (
	"errors"
)

const (
	PreOrderTranverse = 0
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

func (n *TreeNode) String(level int) string {
	//var res string
	//res = fmt.Sprintf("level:%d, str:%s\n", level, n.Value)
	//for _, child := range n.Children {
	//	res += child.String(level + 1)
	//}
	//return res
	return ""
}

func (n *TreeNode) Traversal(method int, res *map[int][]string) error {
	return n.preOrderTraverse(res)
}

func (n *TreeNode) preOrderTraverse(res *map[int][]string) error {
	if res == nil {
		return ErrParamInvalid
	}

	if values, ok := (*res)[n.Level]; ok {
		(*res)[n.Level] = append(values, n.Value)
	} else {
		(*res)[n.Level] = []string{n.Value}
	}

	for _, child := range n.Children {
		child.preOrderTraverse(res)
	}
	return nil
}

func (n *TreeNode) levelOrderTraversal(res *map[int][]string) error {
	return nil
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
	return t.root.String(0)
}

func (t *UrlTree) Traversal(method int) (map[int][]string, error) {
	res := make(map[int][]string)
	err := t.root.Traversal(method, &res)
	return res, err
}
