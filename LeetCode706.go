package main

import "container/list"

type MyMap struct {
	key 	int
	value 	int
}

type MyHashMap struct {
	data []list.List
}

func ConstrutorMap () MyHashMap {
	return MyHashMap{data: make([]list.List , base)}
}

func (this *MyHashMap) Put(key , value int) {
	h := key % base
	for e := this.data[h].Front() ; e != nil ; e = e.Next() {
		if et , _ := e.Value.(MyMap) ; et.key == key {
			e.Value = MyMap{key,value}
			return
		}
	}
	this.data[h].PushBack(MyMap{ key , value })
}

func (this *MyHashMap) Get (key int) int {
	h := key % base
	for e := this.data[h].Front() ; e != nil ; e = e.Next() {
		if et , _ := e.Value.(MyMap) ; et.key == key {
			return et.value
		}
	}
	return -1
}

func (this *MyHashMap) Remove (key int) {
	h := key % base
	for e := this.data[h].Front() ; e != nil ; e = e.Next() {
		if et , _ := e.Value.(MyMap) ; et.key == key {
			this.data[h].Remove(e)
			return
		}
	}
	return
}