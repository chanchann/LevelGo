package memtable

import (
	"bytes"
	"encoding/binary"
)

type ValueType int

const (
	TypeDeletion ValueType = 0
	TypeValue ValueType = 1
)

type InternalKey struct {
	rep []byte
}

func newInternalKey(seq int64, valueType ValueType, key, value []byte) *InternalKey {
	// Format of an entry is concatenation of :
	// 4 : key.size() + 8
	// key bytes : char[key.size()]
	// 8 : seq << 8 | valueType
	// 4 : value.size()
	// value bytes : char[value.size()]
	internalKeySize := len(key) + 8
	valueSize := len(value)
	encodeLen := 4 + internalKeySize + 4 + valueSize 
	buf := make([]byte, encodeLen)

	offset := 0 
	binary.LittleEndian.PutUint32(buf[offset:], uint32(internalKeySize))
	offset += 4
	copy(buf[offset:], key)
	offset += len(key)
	binary.LittleEndian.PutUint64(buf[offset:], (uint64(seq) << 8) | uint64(valueType))
	offset += 8
	binary.LittleEndian.PutUint32(buf[offset:], uint32(valueSize))
	offset += 4
	copy(buf[offset:], value)
	
	return &InternalKey{rep : buf}
}

func (internalKey *InternalKey) userKey() []byte {
	
}

