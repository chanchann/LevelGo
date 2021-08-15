package memtable

import (
	"errors"
	// "github.com/chanchann/LevelGo"
	"github.com/chanchann/LevelGo/skiplist"
)

type MemTable struct {
	table *skiplist.SkipList   //核心就是skiplist
}

func New() *MemTable {
	var memTable MemTable
	memTable.table = skiplist.New(InternalKeyComparator)  // todo
	return &memTable
}

func (memTable *MemTable) NewIterator() levelgo.Iterator {   // todo
	return memTable.table.NewIterator() 
}

func (memTable *MemTable) Add(seq int64, valueType ValueType, key, value []byte) {
	internalKey := newInternalKey(seq, valueType, key, value)
	memTable.table.Insert(internalKey)
}



