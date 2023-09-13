package lc

//LRU -lc 146

//LRUCache lRUCache = new LRUCache(2);
//lRUCache.put(1, 1); // 缓存是 {1=1}
//lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
//lRUCache.get(1);    // 返回 1
//lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
//lRUCache.get(2);    // 返回 -1 (未找到)
//lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
//lRUCache.get(1);    // 返回 -1 (未找到)
//lRUCache.get(3);    // 返回 3
//lRUCache.get(4);    // 返回 4

// 定义头结点是要待删除的 尾节点是最新的呢
type LRUCache struct {
	capacity int
	cache    map[int]*Node
	head     *Node
	tail     *Node
}

type Node struct {
	key  int
	val  int
	pre  *Node
	next *Node
}

func Constructor(capacity int) LRUCache {
	head, tail := &Node{}, &Node{}
	head.next = tail
	tail.pre = head
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		head:     head,
		tail:     tail,
	}
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; ok {
		node := this.cache[key]
		this.deleteNode(node)
		this.addNodeToTail(node)
		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.cache[key]; ok {
		//更新节点value
		this.cache[key].val = value

		//删除节点
		this.deleteNode(this.cache[key])

		//addTail
		this.addNodeToTail(this.cache[key])
	} else {
		//判断容量是否已经满了
		if this.capacity == len(this.cache) {
			//删除头节点
			deleteKey := this.head.next.key
			this.deleteNode(this.head.next)
			delete(this.cache, deleteKey) //删除map
		}
		node := &Node{key: key, val: value}
		this.addNodeToTail(node)
		this.cache[key] = node

	}
}

func (this *LRUCache) addNodeToTail(node *Node) {
	preNode := this.tail.pre
	preNode.next = node
	node.pre = preNode
	node.next = this.tail
	this.tail.pre = node

}

func (this *LRUCache) deleteNode(node *Node) {
	preNode, nextNode := node.pre, node.next
	preNode.next = nextNode
	nextNode.pre = preNode
}
