package data_struct

import (
	"fmt"
	"testing"
)

func TestMakeList(t *testing.T) {
	list, err := MakeList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	if err == nil {
		head := list.head
		if head != nil {
			fmt.Println(head.value)
			head = head.next
		}
	} else {
		t.Failed()
	}
}
