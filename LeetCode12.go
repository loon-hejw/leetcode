package main

import "fmt"

func intToRoman(num int) string {
	Four := []string{"","M","MM","MMM"}
	Three:= []string{"","C","CC","CCC","CD","D","DC","DCC","DCCC","CM"}
	Two  := []string{"","X","XX","XXX","XL","L","LX","LXX","LXXX","XC"}
	One  := []string{"","I","II","III","IV","V","VI","VII","VIII","IX"}
	return Four[num/1000] + Three[(num%1000)/100] + Two[(num%100)/10] + One[num%10]
}
func main() {
	fmt.Println(intToRoman(1994))
}
