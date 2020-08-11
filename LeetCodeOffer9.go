package main

type CQueue struct {
	in []int
	out []int
}


func Constructor() CQueue {
	return CQueue {}
}


func (this *CQueue) AppendTail(value int)  {
	this.in = append(this.in, value)
}

func (this *CQueue) DeleteHead() int {

	if len(this.out) == 0 && len(this.in) == 0 {
		return -1
	}

	if len(this.out) == 0 {
		for len(this.in) > 0 {
			index := len(this.in) - 1
			value := this.in[index]
			this.in = this.in[:index]
			this.out = append(this.out, value)
		}
	}

	index := len(this.out) - 1
	value := this.out[index]
	this.out = this.out[:index]
	return value
}

func main() {
	
}
