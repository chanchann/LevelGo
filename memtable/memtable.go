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

// 将四个参数编码成一个internalKey
func (memTable *MemTable) Add(seq int64, valueType ValueType, key, value []byte) {
	internalKey := newInternalKey(seq, valueType, key, value)
	memTable.table.Insert(internalKey)
}

func (memTable *MemTbale) Get(key []bytes) (bool, []byte, error) {
	lookupKey := LookupKey(key)

	it := memTbale.table.NewIterator()
	it.Seek(lookupKey)
	if it.Valid() {
		internalKey := it.Key().(*InternalKey)
		if UserKeyComparator(key, internalKey.userKey()) == 0 {
			// 判断valueType
			if internalKey.valueType() == TypeValue {
				return true, internalKey.userValue(), nil
			} else {
				return true, nil, erros.New("not found")
			}
		}
	}
	return false, nil, errors.New("not found")
}



