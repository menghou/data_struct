package data_struct

import (
	"fmt"
)

type ThreadedBinaryNode struct {
	data string
	lchild *ThreadedBinaryNode
	rchild *ThreadedBinaryNode
	ltag bool //代表是否是线索
	rtag bool
}

type ThreadedBinaryTree struct {
	root *ThreadedBinaryNode
}

//中序遍历线索化

//传入的必须是还未进行线索化的二叉树
//传入的是树的根节点
func PreOrderInThreading(p *ThreadedBinaryNode, pre *ThreadedBinaryNode) {
	PreOrderInThreading(p.lchild, pre)
}

//平衡二叉树是一个二叉排序树
//每一个节点的左子树和右子树的高度差至多等于1

type AVLTNode struct {
	data   DataType
	h      int
	bf     int
	lchild *AVLTNode
	rchild *AVLTNode
}

func (n *AVLTNode) GetNodeHigh() int {
	if n != nil {
		lchildHigh := n.lchild.GetNodeHigh()
		rchildHigh := n.rchild.GetNodeHigh()
		if lchildHigh > rchildHigh {
			n.h = lchildHigh + 1
		} else {
			n.h = rchildHigh + 1
		}
		return n.h
	} else {
		return 0
	}
}
func (n *AVLTNode) GetBanlanceFactor() int {
	n.bf = n.lchild.GetNodeHigh() - n.rchild.GetNodeHigh()
	return n.bf
}

func (n *AVLTNode) PrintBF() {
	if n != nil {
		n.GetBanlanceFactor()
		fmt.Printf("Data: %d, bf: %d\n", n.data, n.bf)
		n.lchild.PrintBF()
		n.rchild.PrintBF()
	}
}

func (n *AVLTNode) R_Rotate() {
	var l *AVLTNode = n.lchild
	n.lchild = l.rchild
	l.rchild = n
	n = l
}

func (n *AVLTNode) L_Rotate() {
	var r *AVLTNode = n.rchild
	n.rchild = r.lchild
	r.lchild = n
	n = r
}
func (n *AVLTNode) LeftBalance() {
	var l *AVLTNode = n.lchild
	switch l.bf {
	case 1:
		n.bf = 0
		l.bf = 0
		l.R_Rotate()
	case -1:
		var lr *AVLTNode = l.rchild
		switch lr.bf {
		case -1:
			n.bf = 0
			l.bf = 1
		case 0:
			n.bf = 0
			l.bf = 0
		case 1:
			n.bf = 1
			l.bf = 0
		}
		lr.bf = 0
		l.L_Rotate()
		n.R_Rotate()
	}
}

func (n *AVLTNode) RightBalance() {
	var r *AVLTNode = n.rchild
	switch n.bf {
	case -1:
		n.L_Rotate()
		n.bf = 0
		r.bf = 0
	case 1:
		var rl *AVLTNode = r.lchild
		switch rl.bf {
		case -1:
			n.bf = 1
			r.bf = 0
		case 0:
			n.bf = 0
			r.bf = 0
		case 1:
			n.bf = 0
			r.bf = 1
		}
		rl.bf = 0
		r.R_Rotate()
		n.L_Rotate()
	}
}

func (n *AVLTNode) InsertAVL(e DataType, taller *bool) bool {
	if n != nil {
		n = &AVLTNode{data:e}
		n.bf = 0
		*taller = true
	} else {
		if (e == n.data) {
			*taller = false
			return false
		}
		if (e < n.data) {
			if !n.lchild.InsertAVL(e, taller) {
				return false
			}
			if *taller {
				switch n.bf {
				case 1:
					n.LeftBalance()
					*taller = false
				case -1:
					*taller = false
					n.bf = 0
				case 0:
					*taller = true
					n.bf = 1
				}
			}
		} else {
			if !n.rchild.InsertAVL(e, taller) {
				return false
			}
			if *taller {
				switch n.bf {
				case 1:
					n.bf = 0
					*taller = false
				case 0:
					*taller = true
					n.bf = -1
				case -1:
					*taller = false
					n.bf = 0
				}
			}
		}
	}
	return true
}
