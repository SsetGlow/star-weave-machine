package data_struct

import (
	"errors"
)

type list struct {
	size int
	head *node
}

type node struct {
	value int
	next  *node
}

func (l list) Make(values []int) (*list, error) {
	if len(values) <= 0 {
		return nil, errors.New("empty value array")
	}
	var head = node{}
	for index := 0; index < (len(values) - 1); index++ {
		currentValue, nextValue := values[index], values[index+1]
		n := node{currentValue, nil}
		n.next = &node{nextValue, nil}
		if index == 0 {
			head = n
		}
	}
	return &list{1, &head}, nil
}

func (l list) Deconstruct(list list) ([]int, error) {
	if list.size <= 0 {
		return nil, errors.New("empty list")
	}
	res := make([]int, 10, 20)
	head := list.head
	i := 0
	for head != nil {
		res[i] = head.value
		head = head.next
	}
	return res, nil
}
