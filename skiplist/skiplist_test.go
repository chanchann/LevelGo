package skiplist

import (
	"fmt"
	"testing"
	
	"github.com/chanchann/LevelGo/utils"
)

// todo : 把test写的更好

func Test_Insert(t *testing.T) {
	skiplist := New(utils.IntComparator)
	for i:= 0; i < 10; i++ {
		skiplist.Insert(i)
	}
	it := skiplist.NewIterator()
	for it.SeekToFirst(); it.Valid(); it.Next() {
		fmt.Println(it.Key())
	}
	fmt.Println()
	for it.SeekToLast(); it.Valid(); it.Prev() {
		fmt.Println(it.Key())
	}
}