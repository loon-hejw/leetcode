package main

import "fmt"

type HsahSet struct {
	Set []string
	Map map[string]bool
}

func newHashSet() HsahSet {
	return HsahSet{
		Set: []string{},
		Map: map[string]bool{},
	}
}

const startString = "0000"

func openLock(deadends []string, target string) int {

	var (
		beegin = newHashSet()
		end    = newHashSet()
	)

	over := map[string]bool{}
	visited := map[string]bool{}
	count := 1

	for _, v := range deadends {
		over[v] = true
	}

	if over[target] || over[startString] {
		return -1
	}

	// 双端
	beegin.Set = append(beegin.Set, startString)
	beegin.Map[startString] = true
	end.Set = append(end.Set, target)
	end.Map[target] = true

	for len(beegin.Set) > 0 && len(end.Set) > 0 {
		if len(beegin.Set) > len(end.Set) {
			temp := beegin
			beegin = end
			end = temp
		}
		temp := newHashSet()
		for _, v := range beegin.Set {
			for _, str := range getString([]byte(v)) {
				if end.Map[str] {
					return count
				}
				// 去除不能用的
				if over[str] || visited[str] {
					continue
				}
				temp.Set = append(temp.Set, str)
				temp.Map[str] = true
				visited[str] = true
			}
		}
		beegin = temp
		count ++
	}
	return -1
}

func getString(str []byte) (ret []string) {
	for i, b := range str {
		str[i] = b - 1
		if str[i] < '0' {
			str[i] = '9'
		}
		ret = append(ret, string(str))
		str[i] = b + 1
		if str[i] > '9' {
			str[i] = '0'
		}
		ret = append(ret, string(str))
		str[i] = b
	}
	return
}

func main() {

	fmt.Println(openLock([]string{"0201","0101","0102","1212","2002"}, "0202"))
}
