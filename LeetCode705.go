package main

import (
	"container/list"
	"fmt"
)
const base = 769

type MyHashSet struct {
	data []list.List
}

func Constructorset() MyHashSet {
	return MyHashSet{make([]list.List, base)}
}

func (s *MyHashSet) hash(key int) int {
	return key % base
}

func (s *MyHashSet) Add(key int) {
	if !s.Contains(key) {
		h := s.hash(key)
		s.data[h].PushBack(key)
	}
}

func (s *MyHashSet) Remove(key int) {
	h := s.hash(key)
	for e := s.data[h].Front(); e != nil; e = e.Next() {
		if e.Value.(int) == key {
			s.data[h].Remove(e)
		}
	}
}

func (s *MyHashSet) Contains(key int) bool {
	h := s.hash(key)
	for e := s.data[h].Front(); e != nil; e = e.Next() {
		if e.Value.(int) == key {
			return true
		}
	}
	return false
}

func main() {
	MyHashSet := Constructorset()
	MyHashSet.Add(1)   // set = [1]
	MyHashSet.Remove(2)
	MyHashSet.Add(2)      // set = [1, 2]
	fmt.Println(MyHashSet.Contains(2))
	fmt.Println(MyHashSet.Contains(3)) // 返回 False ，（未找到）
	MyHashSet.Add(2)      // set = [1, 2]
	fmt.Println(MyHashSet.Contains(2)) // 返回 True
	MyHashSet.Remove(2)   // set = [1]
	fmt.Println(MyHashSet.Contains(2)) // 返回 False ，（已移除）
}
