package main

// ListNode 链表节点结构体
type ListNode struct {
	Val  int       // 节点值
	Next *ListNode // 指向下一节点的指针
}

// NewListNode 构造函数，创建一个新的链表
func NewListNode(val int) *ListNode {
	return &ListNode{
		Val:  val,
		Next: nil,
	}
}
func main() {

}
