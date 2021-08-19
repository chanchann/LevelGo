package skiplist

// interface{}
// 指定了0个方法的接口值，被称为空接口
// 空接口可保存任何类型的值，所以被用来处理未知类型的值
type Node struct {
	key interface{}   
	next []*Node   
}

// new vs make 
// new(T) 为一个 T 类型新值分配空间并将此空间初始化为 T 的零值，返回的是新值的地址，也就是 T 类型的指针 *T
// make 只能用于 slice,map,channel
// https://sanyuesha.com/2017/07/26/go-make-and-new/
// https://zhuanlan.zhihu.com/p/340988277
func newNode(key interface{}, height int) *Node {
	node := new(Node)
	node.key = key  
	node.next = make([]*Node, height)

	return node
}


// todo : 这里只实现了No-barrier
// todo : memory_order_acquire?
func (node *Node) getNext(level int) *Node {
	return node.next[level]
}

// todo : memory_order_release?
func (node *Node) setNext(level int, x *Node) {
	node.next[level] = x
}