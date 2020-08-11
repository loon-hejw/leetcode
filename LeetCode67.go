package main

func addBinary(a string, b string) string {
	var arr []byte
	if len(a) > len(b) {
		arr = make([]byte,len(a) + 1)
	} else {
		arr = make([]byte,len(b) + 1)
	}
	var sum int
	carry := 0
	for i , j , k := len(a) - 1 , len(b) - 1 , len(arr) - 1 ; i >= 0 || j >= 0 ; {
		sum = carry
		if i >= 0 {
			sum = sum + int(a[i] - '0')
			i = i - 1
		}

		if j >= 0 {
			sum = sum + int(b[j] - '0')
			j = j - 1
		}
		arr[k]= byte(sum % 2) + '0'
		carry = sum / 2
		k --
	}

	if carry != 0 {
		arr[0] = 1 + '0'
		return string(arr)
	}

	return string(arr[1:])
}
func main() {
	addBinary("1010","1011")
}
