package data_struct

import "errors"

type BinaryTree struct {
	head BinaryTreeNode
}

type BinaryTreeNode struct {
	value      int
	leftChild  *BinaryTreeNode
	rightChild *BinaryTreeNode
}

func (t BinaryTree) Make(slice []int) (res *BinaryTree, err error) {
	if len(slice) == 0 {
		err = errors.New("empty slice, not able to make a tree")
		return nil, err
	}
	err = nil

	return res, err
}

func MakeNode(value int, node *BinaryTreeNode) {
	if nil == node {

	} else if nil == node.leftChild {

	} else if nil == node.rightChild {

	}

}
