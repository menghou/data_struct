package data_struct

import (
	"testing"
)

func TestSqList_GetElem(t *testing.T) {
	var test1 *SqList = &SqList{
		head:&SqNode{
			data:5,
			next:NewSqNode(3),
		},
	}
	if err, d0 := test1.GetElem(0); err != nil {
		t.Error(err)
	} else if d0 != 5 {
		t.Errorf("The First Value Should Be %d, Get %d", 5, d0)
	}
	if err, d1 := test1.GetElem(1); err != nil {
		t.Error(err)
	} else if d1 != 3 {
		t.Errorf("The First Value Should Be %d, Get %d", 3, d1)
	}
	if err, d3:= test1.GetElem(2); err == nil {
		t.Errorf("Not Exist 2 Elem, While Get %d", d3)
	}
}

func TestSqList_ListInsert(t *testing.T) {
	test1 := &SqList{head:NewSqNode(3)}
	test1.ListInsert(1,10)
	if err := test1.ListInsert(2, 1); err != nil {
		t.Error(err)
	} else {
		t.Log(test1.String())
	}
}

func TestSqList_ListDelete(t *testing.T) {
	test1 := &SqList{head:NewSqNode(3)}
	test1.ListInsert(1,10)
	if err, r := test1.ListDelete(0); err != nil {
		t.Error(err)
	} else {
		t.Log(test1.String())
		t.Log(r)
	}
}
