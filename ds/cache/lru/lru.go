package lru

/**
基础LRU算法有如下缺点
1.缓存污染 偶尔的批量操作可能导致热点数据删除
*/

//
////LRUCache 记住LRU是双线链表+hash来实现的呢
//type LRUCache struct {
//	capacity int
//	cache    map[int]*Node
//	head     *Node
//	tail     *Node
//}
//
//type Node struct {
//	key   int
//	value int
//	next  *Node
//	pre   *Node
//}
//
//func InitLRUCache(capacity int) *LRUCache {
//	return &LRUCache{
//		capacity: capacity,
//		cache:    make(map[int]*Node),
//		head:     nil,
//		tail:     nil,
//	}
//}

//
//
////
//func (this *LRUCache) Put(key int, value int) {
//	if node, ok := this.cache[key]; ok {
//		node.value = value
//		this.moveToHead(node)
//	} else {
//		newNode := &Node{
//			key:   key,
//			value: value,
//			prev:  nil,
//			next:  nil,
//		}
//		if this.capacity == 0 {
//			delete(this.cache, this.tail.key)
//			this.removeTail()
//		} else {
//			this.capacity--
//		}
//		this.addToHead(newNode)
//		this.cache[key] = newNode
//	}
//}
////
