package main

import (
	"fmt"
	"strings"
)

type ChinaByte struct {
	B bool
	N map[rune]*ChinaByte
}

func(this *ChinaByte) Insert (S string)  {
	if S == "" {
		return
	}
	for _ , v := range S {
		if this.N[v] == nil {
			this.N[v] = &ChinaByte{
				B: false,
				N: map[rune]*ChinaByte{},
			}
		}
		this = this.N[v]
	}
	this.B = true
}

func(this *ChinaByte) Search (S string) string {

	if S == "" {
		return ""
	}
	for i := 0 ; i < len(S) ; i++ {
		t , c := this.SearchPerfix(S[i:])
		if t != nil && t.B && c != ""{
			S = strings.Replace(S ,c , "****" , -1)
			i = i + len(c)
		}
	}
	return S
}

func (this *ChinaByte) SearchPerfix ( word string ) (*ChinaByte , string){

	if word == "" {
		return nil , ""
	}
	t := []rune{}
	for _ , v := range word {
		if this.N[v] == nil {
			return this , string(t)
		}
		this = this.N[v]
		t = append(t , v)
	}
	return this , string(t)
}

func main() {

	CB := &ChinaByte{
		B: false,
		N: map[rune]*ChinaByte{},
	}
	CB.Insert("何嘉伟")
	t := CB.Search("何嘉伟太多的伤感情怀")
	fmt.Println(t)
}
