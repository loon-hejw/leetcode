package main

import "fmt"

func romanToInt(s string) int {
	romanMap := map[byte]int{'O':0,'I':1,'V':5,'X':10,'L':50,'C':100,'D':500,'M':1000}
	sum := 0
	s = s + "O"
	for i := 0 ; i < len(s) - 1 ; i ++{
		t2 := romanMap[s[i+1]]
		t1  := romanMap[s[i]]
		if t2 > t1 {
			sum = sum + (t2 - t1)
			i ++
		} else {
			sum = sum + t1
		}
		t1 = t2 ;
	}
	return sum
}
func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}
