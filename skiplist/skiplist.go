package skiplist

import (
	// "math/rand"
	"sync"
	"github.com/chanchann/LevelGo/utils"  
)
// https://www.google.com/search?q=GOPATH%E8%AE%BE%E7%BD%AE&oq=go&aqs=chrome.1.69i57j69i59l3j69i60l4.2754j0j7&sourceid=chrome&ie=UTF-8

const (
	kMaxHeight = 12
	kBranching = 4
)

// todo : go 也可以实现一些多读并发skiplist
type SkipList struct {
	maxHeight int
	head *Node
	comparator utils.Comparator
	mu sync.RWMutex
}

// go没有构造函数，没有析构函数，需要写个New函数
func New(comp utils.Comparator) *SkipList {
	var skiplist SkipList   
	skiplist.head = newNode(0, kMaxHeight)
	skiplist.maxHeight = 1
	skiplist.comparator = comp  
	return &skiplist
}
