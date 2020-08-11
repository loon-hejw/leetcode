package main

import "fmt"

type Trie struct {
	IsEnd bool
	Next  [26]*Trie
}

/** Initialize your data structure here. */
func Constructor2() Trie {
	return Trie{
		IsEnd: false,
		Next: [26]*Trie{},
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	if word == "" || len(word) == 0 {
		return
	}
	for i := 0 ; i < len(word) ; i ++ {
		n := word[i] - 'a'
		if this.Next[n] == nil {
			this.Next[n] = new(Trie)
		}
		this = this.Next[n]
	}
	this.IsEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	that := this.searchPrefix(word)
	return that != nil && that.IsEnd == true
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	that := this.searchPrefix(prefix)
	return that != nil
}

func (this *Trie) searchPrefix( word string) *Trie {
	for i := 0 ; i < len(word) ; i ++ {
		n := word[i] - 'a'
		if this.Next[n] == nil {
			return nil
		}
		this = this.Next[n]
	}
	return this
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
func main() {
	obj := Constructor2()
	obj.Insert("word")
	param_2 := obj.Search("word")
	param_3 := obj.StartsWith("wo")

	fmt.Println(param_2,param_3)
}