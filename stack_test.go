package data_struct

import "testing"

func TestStack_Push(t *testing.T) {
	stack := &Stack{}
	stack.InitStack()
	stack.Push(1)
	if err, v := stack.GetTop(); err != nil {
		t.Error(err)
	} else if v != 1 {
		t.Errorf("Want 1 Get %d", v)
	} else {
		t.Log("Stack Push OK")
	}
}

func TestStack_ClearStack(t *testing.T) {
	stack := &Stack{}
	stack.InitStack()
	stack.Push(1)
	stack.ClearStack()
	if err, v := stack.GetTop(); err == StackEmpty {
		t.Log("Stack Clear OK")
	} else {
		t.Errorf("Clear Get %d", v)
	}
}

func TestStack_StackLength(t *testing.T) {
	stack := &Stack{}
	stack.InitStack()
	stack.Push(1)
	if l := stack.StackLength(); l != 1 {
		t.Errorf("Want 1 Get %d", l)
	}
	stack.ClearStack()
	if l := stack.StackLength(); l != 0 {
		t.Errorf("Want 1 Get %d", l)
	}
}
