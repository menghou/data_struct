package data_struct

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


