package data_struct

import (
	"testing"
	"fmt"
)

func TestAVLTNode_GetBanlanceFactor(t *testing.T) {
	var a *AVLTNode
	rightAnswer := 0
	if answer := a.GetNodeHigh(); answer != rightAnswer {
		t.Errorf("High Should Be %d, Get: %d", rightAnswer, answer)
	}
}

func TestAVLTNode_GetBanlanceFactor2(t *testing.T) {
	a := AVLTNode{data:1}
	rightAnswer := 1
	if answer := a.GetNodeHigh(); answer != rightAnswer {
		t.Errorf("High Should Be %d, Get: %d", rightAnswer, answer)
	}
}

func TestAVLTNode_GetBanlanceFactor3(t *testing.T) {
	a := &AVLTNode{data:1}
	a.lchild = &AVLTNode{data:2}
	a.rchild = &AVLTNode{data:3}
	rightAnswer := 2
	if answer := a.GetNodeHigh(); answer != rightAnswer {
		t.Errorf("High Should Be %d, Get: %d", rightAnswer, answer)
	}
}

func TestAVLTNode_R_Rotate(t *testing.T) {
	a := &AVLTNode{data:1}
	b := &AVLTNode{data:2}
	a.lchild = b
	a.rchild = &AVLTNode{data:3}
	a.lchild.lchild = &AVLTNode{data:4}
	a.lchild.rchild = &AVLTNode{data:5}
	a.PrintBF()
	a.R_Rotate()
	fmt.Println("")
	a.PrintBF()
	fmt.Println("")
	b.PrintBF()
}

func TestAVLTNode_L_Rotate(t *testing.T) {
	a := &AVLTNode{data:1}
	b := &AVLTNode{data:2}
	c := &AVLTNode{data:3}
	a.lchild = b
	a.rchild = c
	a.rchild.lchild = &AVLTNode{data:4}
	a.rchild.rchild = &AVLTNode{data:5}
	a.PrintBF()
	a.L_Rotate()
	fmt.Println("")
	a.PrintBF()
	fmt.Println("")
	c.PrintBF()
}

func TestAVLTNode_GetBanlanceFactor4(t *testing.T) {
	a := &AVLTNode{data:4}
	a.lchild = &AVLTNode{data:2}
	a.rchild = &AVLTNode{data:6}
	a.lchild.lchild = &AVLTNode{data:1}
	a.lchild.rchild = &AVLTNode{data:3}
	a.rchild.lchild = &AVLTNode{data:5}
	a.rchild.rchild = &AVLTNode{data:9}
	a.rchild.rchild.lchild = &AVLTNode{data:7}
	a.rchild.rchild.rchild = &AVLTNode{data:10}
	a.PrintBF()
}

type AVLTree struct {
	root *AVLTNode
}


