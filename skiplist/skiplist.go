package skiplist

import (
	"math/rand"
	"sync"
	"github.com/chanchann/LevelGo/utils"  
)
// https://www.google.com/search?q=GOPATH%E8%AE%BE%E7%BD%AE&oq=go&aqs=chrome.1.69i57j69i59l3j69i60l4.2754j0j7&sourceid=chrome&ie=UTF-8

const (
	kMaxHeight = 12
	kBranching = 4
)

// todo: need Arena ?  

// todo : go 也可以实现一些多读并发skiplist
// todo : 此处可以用原子操作吗
type SkipList struct {
	maxHeight int
	head *Node
	compare utils.Comparator
	mu sync.RWMutex
}

// go没有构造函数，没有析构函数，需要写个New函数
func New(cmp utils.Comparator) *SkipList {
	var skiplist SkipList   
	skiplist.head = newNode(0, kMaxHeight)
	skiplist.maxHeight = 1
	skiplist.compare = cmp  
	return &skiplist
}

func (list *SkipList) Insert(key interface{}) {
	list.mu.Lock()    // 加写锁
	defer list.mu.Unlock()   // 本函数执行完一定解锁

	_, prev := list.findGreaterOrEqual(key)
	height := list.randomHeight()
	if height > list.maxHeight {
		for i := list.maxHeight; i < height; i++ {
			prev[i] = list.head
		}
		list.maxHeight = height
	}
	x := newNode(key, height)
	for i := 0; i < height; i++ {
		x.setNext(i, prev[i].getNext(i))
		prev[i].setNext(i, x)
	}
}

func (list *SkipList) Contains(key interface{}) bool {
	list.mu.RLock()   // 加读锁
	defer list.mu.RUnlock()
	x, _ := list.findGreaterOrEqual(key)
	if x != nil && list.compare(x.key, key) == 0 {
		return true
	}
	return false
}

func (list *SkipList) NewIterator() *Iterator {
	var it Iterator 
	it.list = list 
	return &it
}

func (list *SkipList) randomHeight() int {
	height := 1
	for height < kMaxHeight && (rand.Intn(kBranching) == 0) {
		height++
	}
	return height
}

func (list *SkipList) findGreaterOrEqual(key interface{}) (*Node, [kMaxHeight]*Node) {
	var prev [kMaxHeight]*Node 
	x := list.head
	level := list.maxHeight - 1
	for true {
		next := x.getNext(level)
		if list.keyIsAfterNode(key, next) {
			x = next  
		} else {
			prev[level] = x
			if level == 0 {
				return next, prev
			} else {
				// switch to next list 
				level--
			}
		}
	}
	return nil, prev
}

func (list *SkipList) findLessThan(key interface{}) *Node {
	x := list.head 
	level := list.maxHeight - 1
	for true {
		next := x.getNext(level)
		if next == nil || list.compare(next.key, key) >= 0 {
			if next == nil || list.compare(next.key, key) >= 0 {
				if level == 0 {
					return x
				} else {
					level--
				}
			} else {
				x = next
			}
		}
	}
	return nil
}

func (list *SkipList) findlast() *Node {
	x := list.head 
	level := list.maxHeight - 1
	for true {
		next := x.getNext(level)
		if next == nil {
			if level == 0 {
				return x
			} else {
				level--
			}
		} else {
			x = next
		}
	}
	return nil
}

func (list *SkipList) keyIsAfterNode(key interface{}, n *Node) bool {
	return (n != nil) && (list.compare(n.key, key) < 0)
}