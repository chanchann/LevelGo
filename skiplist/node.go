package skiplist

// interface{}
// 指定了0个方法的接口值，被称为空接口
// 空接口可保存任何类型的值，所以被用来处理未知类型的值
type Node struct {
	key interface{}   
	next []*Node 
}

func newNode(key interface{}, height int) *Node {
	x := new(Node)
	x.key = key  
	x.next = make([]*Node, height)

	return x
}

func (node *Node) getNext(level int) *Node {
	return node.next[level]
}

func (node *Node) setNext(level int, x *Node) {
	node.next[level] = x
}