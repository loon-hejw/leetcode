package main

import "fmt"

func addStrings(num1 string, num2 string) string {
	sum := 0
	max := len(num1)
	if max < len(num2) {
		max =  len(num2)
	}
	sumChar := make([]byte,max + 1)
	for j , i ,k := len(num1) - 1 , max , len(num2) - 1 ; j >= 0 || k >= 0 ; i -- {
		if j >= 0 {
			sum = sum + int(num1[j] - '0')
			j --
		}
		if  k >= 0 {
			sum = sum + int(num2[k] - '0')
			k --
		}
		sumChar[i] = byte(sum % 10) + '0'
		sum = sum / 10
	}
	if sum > 0 {
		sumChar[0] = byte(sum % 10) + '0'
		return string(sumChar)
	}
	return string(sumChar[1:])
}

func main() {
	k := addStrings("0","0")
	fmt.Println(k)
}
