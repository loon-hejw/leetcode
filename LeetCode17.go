package main

import "fmt"

var ComPhoneMap map[int][]string = map[int][]string{
	1 : []string{} ,
	2 : []string{"a","b","c"} ,
	3 : []string{"d","e","f"} ,
	4 : []string{"g","h","i"} ,
	5 : []string{"j","k","l"} ,
	6 : []string{"m","n","o"} ,
	7 : []string{"p","q","r","s"} ,
	8 : []string{"t","u","v"} ,
	9 : []string{"w","x","y","z"} ,
}
var phoneres []string
func letterCombinations(digits string) []string {
	phoneres = []string{}
	if digits == "" {
		return phoneres
	}
	CombinationsHelper(digits ,"",0)
	return phoneres
}

func CombinationsHelper(digits , res string , index int ) {
	if index == len(digits) {
		phoneres = append(phoneres , res )
	} else {
		t := int(digits[index] - '0')
		temp := ComPhoneMap[t]
		for i := 0 ; i < len(temp) ; i ++ {
			CombinationsHelper(digits ,res + temp[i],index + 1)
		}
	}
}
func main() {
	fmt.Println(letterCombinations("23"))
}
