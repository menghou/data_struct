package data_struct

import (
	"errors"
	"fmt"
	"strings"
)

type DataType int

var (
	overStackError = errors.New("Over Stack")
)

type SqNode struct {
	data DataType
	next *SqNode
}

type SqList struct {
	head *SqNode
}

func (l SqList) String() string {
	tmp := l.head
	resList := make([]string, 0)
	for tmp != nil {
		resList = append(resList, fmt.Sprintf("%d", tmp.data))
		tmp = tmp.next
	}
	return strings.Join(resList, "->")
}

func (l SqList) GetElem(i int) (error, DataType) {
	if l.head == nil {
		return overStackError, 0
	}
	tmp  := l.head
	for j := 0; j < i;  j ++{
		if tmp.next == nil {
			return overStackError, 0
		}
		tmp = tmp.next
	}
	return nil, tmp.data
}

func (l *SqList) ListInsert(i int,d DataType) error {
	if l.head == nil && i != 0 {
		return overStackError
	}
	var tmp *SqNode = l.head
	for j := 0; j < i - 1; j ++ {
		tmp = tmp.next
		if tmp == nil {
			return overStackError
		}
	}
	insertNode := NewSqNode(d)
	insertNode.next = tmp.next
	tmp.next = insertNode
	return nil
}

func (l *SqList) ListDelete(i int) (error,  *SqNode) {
	if i == 0 {
		if l.head != nil {
			tmp := l.head
			l.head = tmp.next
			tmp.next = nil
			return nil, tmp
		} else {
			return overStackError, nil
		}
	}
	var ptmp *SqNode = l.head
	var tmp *SqNode = ptmp.next
	for j := 0; j < i; j ++ {
		if tmp == nil {
			return overStackError, nil
		}
		ptmp = ptmp.next
		tmp = ptmp.next
	}
	fmt.Printf("%d - %d\n",ptmp.data, tmp.data)
	ptmp.next = tmp.next
	tmp.next = nil
	return nil,tmp
}

func NewSqNode(d DataType) *SqNode {
	return &SqNode{data:d}
}
