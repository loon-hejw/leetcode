package main

import (
	"fmt"
	"strconv"
)

func restoreIpAddresses(s string) []string {
	var str []string
	for i := 1 ; i < 4 && i < len(s) -2 ; i ++ {
		for j := i + 1 ; j < i + 4 && j < len(s) - 1 ; j ++ {
			for k := j + 1 ; k < j + 4 &&  k < len(s) ; k ++ {
				str1,str2,str3,str4 := s[0:i] , s[i:j] , s[j:k] , s[k:]
				temp := str1 + "." + str2 + "." + str3 + "." + str4
				if len(temp) == len(s) + 3 && restoreHelper(str1,str2,str3,str4) {
					str = append(str,temp )
				}
			}
		}
	}
	return str
}

func restoreHelper( str1, str2, str3, str4 string) bool {
	int1 , _:= strconv.Atoi(str1)
	int2 , _:= strconv.Atoi(str2)
	int3 , _:= strconv.Atoi(str3)
	int4 , _:= strconv.Atoi(str4)
	if ((len(str1) == 1 && int1 == 0) || (str1[0] > '0' && int1 <= 255 )) &&
		((len(str2) == 1 && int2 == 0) || (str2[0] > '0' && int2 <= 255 )) &&
		((len(str3) == 1 && int3 == 0) || (str3[0] > '0' && int3 <= 255 )) &&
		((len(str4) == 1 && int4 == 0) || (str4[0] > '0' && int4 <= 255 )){
		return true
	}
	return false
}
func main() {
	fmt.Println(restoreIpAddresses("010010"))
}
