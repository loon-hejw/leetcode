package main

import (
	"math"
	"math/bits"
	"strconv"
)

func smallestGoodBase(n string) string {

	num , _ := strconv.Atoi(n)
	max := bits.Len(uint(num)) - 1
	for m := max ; m > 0 ; m -- {
		k := int(math.Pow(float64(num) ,float64(max)))
		mul , sum := 1, 1
		for i := 0 ; i < k ; i ++ {
			mul *= k
			sum += mul
		}
		if sum == num {
			return strconv.Itoa(k)
		}
	}
	return  strconv.Itoa(num - 1)
}

func main() {

}
