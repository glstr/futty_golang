package model

import "errors"

var (
	ErrInvalidNode = errors.New("invalid node")
)

type TreeNode struct {
	Parant   *TreeNode
	Children []*TreeNode
	Value    string
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

type UrlTree struct {
	root *TreeNode
}

func (t *UrlTree) InsertUrl(rawurl string) error {
	uh, err := ParseUrl(rawurl)
	if err != nil {
		return err
	}

	pathSection := uh.GetPathSection()
	for _, section := range pathSection {
		if section == root.Value {
		}
	}

	return
}

func (t *UrlTree) String() string {
	return ""
}
