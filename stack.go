package data_struct

import "errors"

const (
	kStackMaxSize = 200
)

var (
	StackEmpty = errors.New("Stack Empty")
	StackFull = errors.New("Stack Full")
)

//需要考虑栈满该怎么办
type Stack struct {
	data []DataType
	top int
}
//初始化操作，简历一个空栈
func (s *Stack) InitStack() {
	s.data = make([]DataType, kStackMaxSize)
	s.ClearStack()
}

func (s *Stack) ClearStack() {
	s.top = -1
}

func (s *Stack) StackEmpty() bool {
	return s.top == -1
}

func (s *Stack) GetTop() (error, DataType) {
	if !s.StackEmpty() {
		return nil, s.data[s.top]
	}
	return StackEmpty, 0
}

func (s *Stack) Push(a DataType) error {
	if s.top == kStackMaxSize -1 {
		return StackFull
	} else {
		s.top += 1
		s.data[s.top] = a
		return nil
	}
}

func (s *Stack) Pop() (error, DataType) {
	if s.StackEmpty() {
		return StackEmpty, 0
	}
	s.top -= 1
	return nil, s.data[s.top-1]
}

func (s *Stack) StackLength() int {
	return s.top + 1
}

type SqStackNode struct {
	data DataType
	next *SqStackNode
}

type SqStack struct {
	head *SqStackNode
}
