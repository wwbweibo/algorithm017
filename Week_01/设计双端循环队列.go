type LinkedListNode struct {
	Value int
	Prev *LinkedListNode
	Next *LinkedListNode
}

type MyCircularDeque struct {
	Length int
	Capacity int
	Head *LinkedListNode
	Tail *LinkedListNode
}


/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	// 初始化时，头节点Next直接指向尾节点
	// 尾节点的Prev指向Head
	headNode := LinkedListNode{
		Value: -1,
	}
	tailNode := LinkedListNode{
		Value: -1,
	}
	headNode.Next = &tailNode
	tailNode.Prev = &headNode
	return MyCircularDeque{
		Capacity: k,
		Length: 0,
		Head: &headNode,
		Tail: &tailNode,
	}
}


/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}

	this.Length += 1
	node := LinkedListNode{
		Next: this.Head,
		Value: -1,
	}
	this.Head.Value = value
	this.Head.Prev = &node
	this.Head = &node
	return true
}


/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	this.Length += 1
	node := LinkedListNode{
		Prev: this.Tail,
		Value: -1,
	}
	this.Tail.Value = value
	this.Tail.Next = &node
	this.Tail = &node
	return true
}


/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.Length -= 1

	this.Head.Next.Next.Prev = this.Head
	this.Head.Next = this.Head.Next.Next
	return true
}


/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.Length -= 1
	this.Tail.Prev.Prev.Next = this.Tail
	this.Tail.Prev = this.Tail.Prev.Prev
	return true
}


/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	return this.Head.Next.Value
}


/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	return this.Tail.Prev.Value
}


/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	return this.Length == 0
}


/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	return this.Length == this.Capacity
}