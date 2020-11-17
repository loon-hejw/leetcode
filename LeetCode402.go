package main

import "fmt"

func removeKdigits(num string, k int) string {

	if num == "" {
		return num
	}

	List := []byte{}
	for i := 0 ; i < len(num) ; i++ {
		for len(List) > 0 && List[len(List) - 1] > num[i] && k > 0 {
			List = List[:len(List) - 1]
			k = k - 1
		}
		List = append(List , num[i])
	}

	for k > 0 && len(List) > 0 {
		List = List[:len(List) - 1]
		k --
	}
	for len(List) > 0 && List[0] == '0' {
		List = List[1:]
	}

	if len(List) == 0 {
		return "0"
	}
	return string(List)
}

func main() {
	fmt.Println(removeKdigits("10200" , 1))
}
