package main

import "fmt"

/**
 * 自己写的测试 平衡二叉树 ～～～
 * 还有点小小的问题 大致就是这样的结构
 */
type AvlTreeNode struct {
 	Key int
	Height int
 	Left *AvlTreeNode
 	Right *AvlTreeNode
}

// 获取树的高度
func (this *AvlTreeNode) AvlTreeHeight() int {
	if nil == this {
		return 0
	}
	return this.Height
}

// 获取最大值
func getAvlMax ( y , x int) int {
	if y > x {
		return y
	}
	return x
}

// 获取平衡因子
func (this *AvlTreeNode) getBalance () int {
	if nil == this {
		return 0
	}
	return  this.Left.AvlTreeHeight() - this.Right.AvlTreeHeight()
}

// 右旋
func (this *AvlTreeNode) rightRotate () *AvlTreeNode {
	x := this.Left
	t2 := x.Right
	// 右旋转
	x.Right = this
	this.Left = t2
	// 更新高度
	this.Height = getAvlMax( this.Left.Height , this.Right.Height ) + 1
	x.Height    = getAvlMax( x.Left.Height    , x.Right.Height    ) + 1
	return x
}

// 左旋
func (this *AvlTreeNode) leftRotate() *AvlTreeNode {
	y := this.Right
	t2 := y.Left
	// 左旋转
	y.Left = this
	this.Right = t2

	this.Height = getAvlMax( this.Left.Height , this.Right.Height ) + 1
	y.Height    = getAvlMax( y.Left.Height    , y.Right.Height    ) + 1
	return  y
}

func (this *AvlTreeNode) Inset(key int) *AvlTreeNode {
	if nil == this {
		return &AvlTreeNode{
			Key: key,
			Height: 1,
		}
	}
	if key < this.Key {
		this.Left = this.Left.Inset(key)
	} else if key > this.Key {
		this.Right = this.Right.Inset(key)
	} else {
		return this
	}

	balance := this.getBalance()
	if balance > 1 && key < this.Left.Key {
		return this.rightRotate()
	}
	if balance < -1 && key > this.Right.Key {
		return this.leftRotate()
	}
	if balance > 1 && key > this.Left.Key {
		this.Left = this.leftRotate()
		return this.rightRotate()
	}
	if balance < -1 && key > this.Right.Key {
		this.Right = this.rightRotate()
		return this.leftRotate()
	}
	return this
}

func (this *AvlTreeNode) preOrder (){
	if nil != this {
		fmt.Print(this.Key)
		fmt.Print(" ")
		this.Left.preOrder()
		this.Right.preOrder()
	}
}
func main() {
	var root AvlTreeNode
	//root := &AvlTreeNode{Key : 10 ,Height : 1}
	root.Inset(10)
	root.Inset(20)
	root.Inset(30)
	root.Inset(40)
	root.Inset(50)
	root.Inset(25)
	root.preOrder()
}