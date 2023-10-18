package data_struct

import "errors"

type BinaryTree struct {
	Size  int
	Depth int
	head  BinaryTreeNode
}

type BinaryTreeNode struct {
	value      int
	leftChild  *BinaryTreeNode
	rightChild *BinaryTreeNode
}

func MakeTree(slice []int) (res *BinaryTree, err error) {
	if len(slice) == 0 {
		err = errors.New("empty slice, not able to make a tree")
		return nil, err
	}
	err = nil
	head := BinaryTreeNode{value: slice[0], leftChild: nil, rightChild: nil}
	for _, i := range slice {
		MakeNode(i, &head)
	}
	res = &BinaryTree{Size: len(slice), Depth: 0, head: head}
	return res, err
}

func MakeNode(value int, node *BinaryTreeNode) {
	if nil == node {

	} else if nil == node.leftChild {

	} else if nil == node.rightChild {

	}
}

func (t BinaryTree) PreorderTraverse() (res *[]int) {
	return res
}

func (t BinaryTree) InorderTraverse() (res *[]int) {
	return res
}

func (t BinaryTree) PostorderTraverse() (res *[]int) {
	return res
}
