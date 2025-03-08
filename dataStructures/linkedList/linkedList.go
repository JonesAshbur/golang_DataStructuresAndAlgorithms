package main

import "fmt"

// ListNode 单链表节点结构体
type ListNode struct {
	Val  int       // 节点值
	Next *ListNode // 指向下一节点的指针
}

// DoublyListNode 双向链表
type DoublyListNode struct {
	Val  int             // 节点值
	Next *DoublyListNode // 指向后继节点的指针
	Prev *DoublyListNode // 指向前驱节点的指针
}

// NewDoublyListNode 双向链表初始化
func NewDoublyListNode(val int) *DoublyListNode {
	return &DoublyListNode{
		Val:  val,
		Next: nil,
		Prev: nil,
	}
}

// NewListNode 构造函数，创建一个新的链表
func NewListNode(val int) *ListNode {
	return &ListNode{
		Val:  val,
		Next: nil,
	}
}

// n之后插入节点p，时间复杂度为O（1）
func InsertNode(n *ListNode, p *ListNode) {
	tag := n.Next
	p.Next = tag
	n.Next = p
}

// 删除链表的节点 n 之后的首个节点，时间复杂度O（1）
func removeItem(n *ListNode) {
	if n.Next == nil {
		return
	}
	// n-> P -> m
	p := n.Next
	m := p.Next
	n.Next = m
	//n.Next = n.Next.Next
}

// 访问第index个节点，时间复杂度O(n)
func access(head *ListNode, index int) *ListNode {
	for i := 0; i < index; i++ {
		if head == nil {
			return nil
		}
		head = head.Next
	}
	return head
}

// 查找值为target的第一个节点
func findNode(head *ListNode, target int) int {
	index := 0
	for head != nil {
		if head.Val == target {
			return index
		}
		head = head.Next
		index++
	}
	return -1
}
func main() {
	//初始化链表数据
	num0 := NewListNode(0)
	num1 := NewListNode(1)
	num2 := NewListNode(2)
	num3 := NewListNode(3)
	num4 := NewListNode(4)
	//构建节点间的链接顺序
	num0.Next = num1
	num1.Next = num2
	num2.Next = num3
	num3.Next = num4
	fmt.Println(*num0, *(num0.Next), *(num1.Next), *(num2.Next), *(num3.Next))
	//插入节点
	num5 := NewListNode(5)
	InsertNode(num0, num5)
	fmt.Println(*num0, *(num0.Next), *(num5.Next), *(num1.Next), *(num2.Next), *(num3.Next))
	//删除节点num0后的第一个节点
	removeItem(num0)
	fmt.Println(*num0, *(num0.Next), *(num1.Next), *(num2.Next), *(num3.Next))
	//访问节点
	res1 := access(num0, 3)
	fmt.Println("第三个节点的数据：", (*res1).Val)
	//查找结点
	res2 := findNode(num0, 1)
	fmt.Println("目标值index=", res2)
}
