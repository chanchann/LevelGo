package skiplist

import (
	"math/rand"
	"sync"
	// ""  // todo : 如何import utils
)

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

func New(comp utils.Comparator) *SkipList {
	var skiplist SkipList   
	skiplist.head = newNode(0, kMaxHeight)
	skiplist.maxHeight = 1
	skiplist.comparator = comp  
	return &skiplist
}
