package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool  { return cnt[h.IntSlice[i]] > cnt[h.IntSlice[j]] }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{}   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
func (h *hp) push(v int)         { heap.Push(h, v) }
func (h *hp) pop() int           { return heap.Pop(h).(int) }

func reorganizeString(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}

	cnt = [26]int{}
	maxCnt := 0
	for _, ch := range s {
		ch -= 'a'
		cnt[ch]++
		if cnt[ch] > maxCnt {
			maxCnt = cnt[ch]
		}
	}
	if maxCnt > (n+1)/2 {
		return ""
	}

	h := &hp{}
	for i, c := range cnt[:] {
		if c > 0 {
			h.IntSlice = append(h.IntSlice, i)
		}
	}
	heap.Init(h)

	ans := make([]byte, 0, n)
	for len(h.IntSlice) > 1 {
		i, j := h.pop(), h.pop()
		ans = append(ans, byte('a'+i), byte('a'+j))
		if cnt[i]--; cnt[i] > 0 {
			h.push(i)
		}
		if cnt[j]--; cnt[j] > 0 {
			h.push(j)
		}
	}
	if len(h.IntSlice) > 0 {
		ans = append(ans, byte('a'+h.IntSlice[0]))
	}
	return string(ans)
}

func main() {
	fmt.Println(reorganizeString("tndsewnllhrtwsvxenkscbivijfqnysamckzoyfnapuotmdexzkkrpmppttficzerdndssuveompqkemtbwbodrhwsfpbmkafpwyedpcowruntvymxtyyejqtajkcjakghtdwmuygecjncxzcxezgecrxonnszmqmecgvqqkdagvaaucewelchsmebikscciegzoiamovdojrmmwgbxeygibxxltemfgpogjkhobmhwquizuwvhfaiavsxhiknysdghcawcrphaykyashchyomklvghkyabxatmrkmrfsppfhgrwywtlxebgzmevefcqquvhvgounldxkdzndwybxhtycmlybhaaqvodntsvfhwcuhvuccwcsxelafyzushjhfyklvghpfvknprfouevsxmcuhiiiewcluehpmzrjzffnrptwbuhnyahrbzqvirvmffbxvrmynfcnupnukayjghpusewdwrbkhvjnveuiionefmnfxao"))
}