package data_struct

import (
	"errors"
)

type LinkedList struct {
	size int
	head *node
}

type node struct {
	value int
	next  *node
}

func MakeList(slice []int) (*LinkedList, error) {
	if len(slice) <= 0 {
		return nil, errors.New("empty value array")
	}
	var head = node{slice[0], nil}
	res := &LinkedList{len(slice), &head}
	for index := 1; index < len(slice); index++ {
		head.next = &node{slice[index], nil}
		head = *head.next
	}
	return res, nil
}
