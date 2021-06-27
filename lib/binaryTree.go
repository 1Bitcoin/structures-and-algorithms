package lib

import (
	"errors"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) Insert(value int) error {
	if t == nil {
		return errors.New("Tree is nil")
	}

	if t.Val == value {
		return errors.New("This node value already exists")
	}

	if t.Val > value {
		if t.Left == nil {
			t.Left = &TreeNode{Val: value}
			return nil
		}
		return t.Left.Insert(value)
	}

	if t.Val < value {
		if t.Right == nil {
			t.Right = &TreeNode{Val: value}
			return nil
		}
		return t.Right.Insert(value)
	}
	return nil
}

//Find finds the treenode for the given node val
func (t *TreeNode) Find(value int) (TreeNode, bool) {
	if t == nil {
		return TreeNode{}, false
	}

	switch {
	case value == t.Val:
		return *t, true
	case value < t.Val:
		return t.Left.Find(value)
	default:
		return t.Right.Find(value)
	}
}

func (t *TreeNode) FindMin() int {
	if t.Left == nil {
		return t.Val
	}
	return t.Left.FindMin()
}

func (t *TreeNode) FindMax() int {
	if t.Right == nil {
		return t.Val
	}
	return t.Right.FindMax()
}

func (t *TreeNode) PrintInorder() {
	if t == nil {
		return
	}

	t.Left.PrintInorder()
	fmt.Print(t.Val, " ")
	t.Right.PrintInorder()
}
