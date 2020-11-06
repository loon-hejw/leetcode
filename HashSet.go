package main

import "fmt"

/*
  实现 set 集合

  原理使用原有的HashMap的数据结构完成

  能够在o(1)的时间复杂度判断出是否存在集合中

  能够在o(1)的时间复杂度集合的长度
 */

type HashSet struct {
	length int
	set    map[interface{}]bool
	val    []interface{}
}

func (this *HashSet) GetLenght() int {
	return this.length
}

func (this *HashSet) Set(val interface{}) {
	if !this.set[val] {
		this.length = this.length + 1
		this.val = append(this.val , val)
	}
}

func (this *HashSet) Get () interface{} {
	var r interface{}
	if this.length > 0 {
		r = this.val[0]
		this.val = this.val[1:]
		this.length = this.length - 1
		return r
	}
	return r
}

func (this *HashSet) inset (val interface{}) bool {
	return this.set[val]
}

func getHashSet() *HashSet {
	return &HashSet{
		length : 0,
		set    : map[interface{}]bool{},
		val    : []interface{}{},
	}
}

func main() {
	mySet := getHashSet()
	fmt.Println(mySet.inset("1111111"))
	mySet.Set("11111111")
	fmt.Println(mySet.GetLenght())
	fmt.Println(mySet.Get())
}
