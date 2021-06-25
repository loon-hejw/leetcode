package main

import (
	"fmt"
	"math/bits"
	"math/rand"
	"runtime"
	"time"
)

func readBinaryWatch(turnedOn int) []string {
	ans := make([]string, 0)
	for h := uint(0); h < 12; h++ {
		for m := uint(0); m < 60; m++ {
			if bits.OnesCount(h)+bits.OnesCount(m) == turnedOn {
				ans = append(ans, fmt.Sprintf("%d:%02d", h, m))
			}
		}
	}
	return ans
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println(runtime.NumCPU())
	fmt.Println(readBinaryWatch(1))
}
