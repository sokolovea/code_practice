package main

import "fmt"

type KeyValue struct {
	Key   int
	Value int
}

type DoubleLinkedElement struct {
	Data KeyValue
	Prev *DoubleLinkedElement
	Next *DoubleLinkedElement
}

type DoubleLinkedList struct {
	First    *DoubleLinkedElement
	Last     *DoubleLinkedElement
	Length   int
	Capacity int
}

func (this *DoubleLinkedList) Add(key int, value int) *int {
	var deleteKey *int = nil
	if this.Length >= this.Capacity {
		deleteKey = &this.First.Data.Key
		this.DeleteFirst()
	}
	newElem := new(DoubleLinkedElement)
	newElem.Data.Key = key
	newElem.Data.Value = value
	if this.Last == nil {
		this.Last = newElem
		this.First = this.Last
	} else {
		this.Last.Next = newElem
		newElem.Prev = this.Last
		this.Last = newElem
	}
	this.Length++
	return deleteKey
}

func (this *DoubleLinkedList) DeleteFirst() {
	if this.First != nil {
		this.First = this.First.Next
		if this.First != nil {
			this.First.Prev = nil
		} else {
			this.Last = nil
		}
		this.Length -= 1
	}
}

func (this *DoubleLinkedList) updateTTL(node *DoubleLinkedElement) {
	if node == this.Last {
		return
	}
	// cut
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}
	if node.Prev == nil {
		this.First = node.Next
	} else {
		node.Prev.Next = node.Next
	}
	// move to the last
	node.Next = nil
	node.Prev = this.Last
	if this.Last != nil {
		this.Last.Next = node
	}
	this.Last = node
}

type LRUCache struct {
	KeyMap    map[int]*DoubleLinkedElement
	ValueList DoubleLinkedList
	Capacity  int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		KeyMap: make(map[int]*DoubleLinkedElement),
		ValueList: DoubleLinkedList{
			Capacity: capacity,
		},
		Capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	elemPtr := this.KeyMap[key]
	if elemPtr == nil {
		return -1
	}
	value := elemPtr.Data.Value
	this.ValueList.updateTTL(elemPtr)
	return value
}

func (this *LRUCache) Put(key int, value int) {
	elemPtr := this.KeyMap[key]
	if elemPtr != nil {
		elemPtr.Data.Value = value
		this.ValueList.updateTTL(elemPtr)
	} else {
		deleteKey := this.ValueList.Add(key, value)
		if deleteKey != nil {
			delete(this.KeyMap, *deleteKey)
		}
	}
	this.KeyMap[key] = this.ValueList.Last
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
