package main

import "fmt"

func multiply(num1 string, num2 string) string {

	if num1 == "0" || num2 == "0" {
		return "0"
	}
	sum := make([]int,len(num1) + len(num2))
	for i := len(num1) - 1 ; i >= 0 ; i -- {
		for j := len(num2) - 1 ; j >= 0 ; j -- {
			sum[i + j + 1] += int(num1[i] - '0') * int(num2[j] - '0')
		}
	}
	str := make([]byte,len(num1) + len(num2))
	for i := len(num1) + len(num2) - 1 ; i > 0 ; i -- {
		sum[i - 1] = sum[i] / 10 + sum[i - 1]
		sum[i] = sum[i] % 10
		str[i - 1] = byte(sum[i - 1] + '0')
		str[i] = byte(sum[i] + '0')
	}
	if str[0] == '0' {
		return string(str[1:])
	}
	return string(str)
}
func main() {
	fmt.Println(multiply("123","456"))
}
