package main

type doublyLinkedNode struct{
	prev, next *doublyLinkedNode
	key, val int
}
type LRUCache struct {
	len, cap int
	first, last *doublyLinkedNode
	m map[int]*doublyLinkedNode
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap	: capacity,
		m	: make(map[int]*doublyLinkedNode, capacity),
	}
}

func (c *LRUCache) Get(key int) int {
	// 存在则返回
	if node, ok := c.m[key]; ok{
		c.moveToFirst(node)
		return node.val
	}
	return -1
}

func (c *LRUCache) Put(key int, value int)  {
	node, ok := c.m[key]
	if ok {// 更新
		node.val = value
		c.moveToFirst(node) //如果不放到前面，那么后面的操作可能会将它错误淘汰
	}else{
		// 插入 淘汰旧数据
		if c.len==c.cap{
			delete(c.m, c.last.key)
			c.removeLast()
		}else{
			c.len++
		}
		node = &doublyLinkedNode{
			key: key,
			val: value,
		}
		c.m[key] = node
		c.insertToFirst(node)
	}
}

func (c *LRUCache)removeLast(){
	// 删除末尾节点：更新倒数第二个节点指针
	if c.last.prev!= nil{// 非空链表
		c.last.prev.next = nil
	}else{
		c.first = nil
	}
	c.last = c.last.prev
}
func (c *LRUCache)moveToFirst(node *doublyLinkedNode){
	// 先删除再插入
	switch node{
		case c.first:
			return
		case c.last:
			c.removeLast()
		default:
			// 删除节点：更新前后节点指针
			node.prev.next = node.next
			node.next.prev = node.prev
	}
	c.insertToFirst(node)
}

func (c *LRUCache)insertToFirst(node *doublyLinkedNode){
	// 异常（空链表）
	if c.last==nil{
		c.last = node
	}else{
		// 双向指针
		node.next = c.first
		c.first.prev = node
	}
	// 指定头节点
	c.first = node
}
