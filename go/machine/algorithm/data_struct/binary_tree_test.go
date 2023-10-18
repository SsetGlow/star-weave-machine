package data_struct

import (
	"fmt"
	"testing"
)

func TestMakeTree(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	tree, err := MakeTree(slice)
	if nil == err {
		fmt.Println(tree)
	} else {
		t.Error("fail to build")
	}
}
