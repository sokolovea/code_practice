package main

import "fmt"

type KeyValue struct {
	Key   int
	Value int
	Valid bool
}

type RingBuffer struct {
	buffer   []KeyValue
	curIndex int
}

func (this *RingBuffer) Push(value KeyValue) {
	var newElemIndex int = (this.curIndex + 1) % len(this.buffer)
	this.buffer[newElemIndex] = value
	this.curIndex = newElemIndex
}

func (this *RingBuffer) AfterTop() KeyValue {
	return this.buffer[(this.curIndex+1)%len(this.buffer)]
}

func (this *RingBuffer) topPtr() *KeyValue {
	return &this.buffer[this.curIndex]
}

func (this *RingBuffer) updateTTL(
	updatedElementPtr *KeyValue) (*KeyValue, *KeyValue) {
	if &this.buffer[this.curIndex] != updatedElementPtr {
		*updatedElementPtr, this.buffer[this.curIndex] =
			this.buffer[this.curIndex], *updatedElementPtr
	}
	return &(this.buffer[this.curIndex]), updatedElementPtr
}

func (this *RingBuffer) Get(index int) KeyValue {
	return this.buffer[this.curIndex]
}

type LRUCache struct {
	KeyMap          map[int]*KeyValue
	ValueRingBuffer RingBuffer
	Capacity        int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		KeyMap: make(map[int]*KeyValue),
		ValueRingBuffer: RingBuffer{
			buffer:   make([]KeyValue, capacity),
			curIndex: capacity - 1,
		},
		Capacity: capacity,
	}
}

func (this *LRUCache) updateTTL(elemPtr *KeyValue) {
	exchangedKeyValuePtr1, exchangedKeyValuePtr2 :=
		this.ValueRingBuffer.updateTTL(elemPtr)
	if exchangedKeyValuePtr1 != exchangedKeyValuePtr2 {
		this.KeyMap[(*exchangedKeyValuePtr1).Key] = exchangedKeyValuePtr1
		this.KeyMap[(*exchangedKeyValuePtr2).Key] = exchangedKeyValuePtr2
	}
}

func (this *LRUCache) Get(key int) int {
	elemPtr := this.KeyMap[key]
	if elemPtr == nil {
		return -1
	}
	value := *elemPtr
	this.updateTTL(elemPtr)
	return value.Value
}

func (this *LRUCache) Put(key int, value int) {
	elemPtr := this.KeyMap[key]
	if elemPtr == nil {
		deleteKeyValue := this.ValueRingBuffer.AfterTop()
		//maybe not init if RingBuffer partially filled
		if deleteKeyValue.Valid {
			delete(this.KeyMap, deleteKeyValue.Key)
		}
		this.ValueRingBuffer.Push(KeyValue{Key: key, Value: value, Valid: true})
	} else {
		elemPtr.Value = value
		this.updateTTL(elemPtr)
	}
	this.KeyMap[key] = this.ValueRingBuffer.topPtr()
}

func main() {
	lRUCache := Constructor(2)
	lRUCache.Put(1, 1)           // cache is {1=1}
	lRUCache.Put(2, 2)           // cache is {1=1, 2=2}
	fmt.Println(lRUCache.Get(1)) // return 1
	lRUCache.Put(3, 3)           // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
	fmt.Println(lRUCache.Get(2)) // returns -1 (not found)
	lRUCache.Put(4, 4)           // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
	fmt.Println(lRUCache.Get(1)) // return -1 (not found)
	fmt.Println(lRUCache.Get(3)) // return 3
	fmt.Println(lRUCache.Get(4)) // return 4
}
